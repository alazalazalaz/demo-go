为了解决各个项目对不同翻译渠道的需求，平台开发了统一的翻译服务。

目前支持将Google、MicroSoft、Amazon作为翻译服务提供源。并且内置翻译缓存，以减少对翻译服务的重复请求导致的成本浪费。

## 接入须知

翻译服务是开箱即用的，您不需要做任何的接入准备即可使用。

默认情况下，您的翻译请求有40%几率会使用Microsoft渠道进行翻译、0.1%几率使用Amazon、59.9%使用Google翻译。

对于部分Amazon翻译返回的错误结果（小概率、取决于翻译服务商当下服务稳定性以及输入是否合法等因素），我们会使用MicroSoft渠道进行重新翻译。

另外，我们在翻译服务中增加了服务切换功能，对于各服务商当前的服务质量进行跟踪。当某服务商当下处于服务异常时，会优先使用其他服务商的翻译服务，确保平台服务稳定可用。

## 服务地址

翻译服务的地址请通过`Name`服务获取，服务名为`translate`。

## API文档
请查阅[Translate服务的API文档](https://wiki.tap4fun.com/display/TGS/Translate-API)

## 免翻译文本处理

针对Emoji表情、特殊内容等需要禁止的翻译内容，请使用`translate`属性值为`no`的HTML标签进行包裹。参见 [HTML标准文档](https://developer.mozilla.org/zh-CN/docs/Web/HTML/Global_attributes/translate)

例如`我爱<span translate="no">北京</span>天安门`中的`北京`就可以不会被翻译。

详见 [Google](https://cloud.google.com/translate/troubleshooting) 的文档以及 [Microsoft](https://docs.microsoft.com/en-us/azure/cognitive-services/translator/prevent-translation) 的文档

## 翻译内容对照

由于我们的翻译服务商有多家，如果您对翻译结果有疑问，可以对照这几家的官方翻译检查结果：

- [Google](https://www.google.com/translate)
- [Microsoft](https://www.bing.com/translator)
- [AWS](https://console.aws.amazon.com/translate/home) （需要登陆）

## 流程图

### 项目启动流程图

```mermaid
graph LR
    begin([项目启动])-->A["重载OnStart()方法"]
    subgraph one ["重载OnStart()方法"]
    A-->B["初始化AWS相关"]
    A-->C["硬编码初始化各个渠道语种简称"]
    B & C-->D["读取配置文件更新服务商对象信息等"]
    A-.->E[新启groutine]
    E-.->|每隔5分钟|D
    D-->F[注册转换函数,稍后使用]
    end
    F-->G["设置路由(thrift代码生成)"]
    G-->H["读取json_patch_paths配置文件"]
    H-->I(["重载OnStart()完成"])
```


### translate接口流程图

```mermaid
graph TB
    begin(["开始处理请求translate或者google/translate"])-->A{校验权限}
    subgraph one
    A-->|失败|over1[结束]
    A-->|成功|B{目标语种}
    end
    B-->|单目标|C[单目标翻译]
    B-->|单目标|D[多目标翻译]
    C & D -->CD1[选择服务商]
    CD1-->CD2{命中缓存}
    CD2-.->CD3[(redis和db)]
    CD2-->|否|E[选择渠道商]
    CD2-->|是|over2[结束]

    E-->F1[微软翻译]
    E-->F2[亚马逊翻译]
    E-->F3[谷歌翻译]

    F1-->F11[源和目标语言进行一次转换<br>谷歌转为微软]
    F2-->F21[源和目标语言进行一次转换<br>谷歌转为亚马逊]
    F3-->F31[无需转换]
    
    F11 & F21 & F31-->G["Translate()"]
    G-->HA["MicrosoftTranslate()"]
    G-->HB["AmazonTranslate()"]
    G-->HC["GoogleTranslate()"]
    subgraph oneA["微软"]
        HA-->HAA{命中缓存}
        
        HAA-->|是|HAD[语言文本转换]
        HAA-->|否|HAC{curl微软}
        HAC-->|是|HAF[统计字符,写入redis]
        HAC-->|否|HAE[打点+1,redis统计]
        HAE-->over3[结束]
        
        HAF-->HAD
    end
    subgraph oneB["亚马逊"]
        HB-->HBA{命中缓存}
        
        HBA-->|是|HBD[语言文本转换]
        HBA-->|否|HBC{curl微软}
        HBC-->|是|HBF[统计字符,写入redis]
        HBC-->|否|HBE[打点+1,redis统计]
        HBE-->over4[结束]
        
        HBF-->HBD
    end
    subgraph oneC["谷歌"]
        HC-->HCA{命中缓存}
        
        HCA-->|是|HCD[无需转换]
        HCA-->|否|HCC{curl微软}
        HCC-->|是|HCF[统计字符,写入redis]
        HCC-->|否|HCE[打点+1,redis统计]
        HCE-->over5[结束]
        
        HCF-->HCD
    end
    
    
    HAD & HBD & HCD -->I[更新缓存]
    I-->over6([返回成功])
```