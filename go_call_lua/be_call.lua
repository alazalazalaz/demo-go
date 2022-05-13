function GetString()
    print "hello world"
end

function getBigger(a, b)
    if a > b then
        print(a)
    else
        print(b)
    end
end

function printString(str, str2)
    print("prepare print:" .. str)
    print("prepare print2:" .. str2)
end

function regText(str)
    newlog = string.gsub( str ,".*%[DataName%]:([^,]+),%[Data%]:(.*)\n","{\"data\":%2,\"topic\":\"%1\"}")
    print(newlog)
end



--printString("im lua \\u200f")