package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
)

func main() {
	azureOpenAIKey := os.Getenv("AZURE_OPENAI_KEY")
	//modelDeploymentID = deployment name, if model name and deployment name do not match change this value to name chosen when you deployed the model.
	modelDeploymentID := "gpt-35-turbo"

	// Ex: "https://<your-azure-openai-host>.openai.azure.com"
	azureOpenAIEndpoint := os.Getenv("AZURE_OPENAI_ENDPOINT")

	if azureOpenAIKey == "" || modelDeploymentID == "" || azureOpenAIEndpoint == "" {
		fmt.Fprintf(os.Stderr, "Skipping example, environment variables missing\n")
		return
	}

	keyCredential, err := azopenai.NewKeyCredential(azureOpenAIKey)

	if err != nil {
		//  TODO: Update the following line with your application specific error handling logic
		log.Fatalf("ERROR: %s", err)
	}

	client, err := azopenai.NewClientWithKeyCredential(azureOpenAIEndpoint, keyCredential, nil)

	if err != nil {
		//  TODO: Update the following line with your application specific error handling logic
		log.Fatalf("ERROR: %s", err)
	}

	// This is a conversation in progress.
	// NOTE: all messages, regardless of role, count against token usage for this API.
	messages := []azopenai.ChatMessage{
		// You set the tone and rules of the conversation with a prompt as the system role.
		{Role: to.Ptr(azopenai.ChatRoleSystem), Content: to.Ptr("You are a helpful assistant.")},

		// The user asks a question
		{Role: to.Ptr(azopenai.ChatRoleUser), Content: to.Ptr("Does Azure OpenAI support customer managed keys?")},

		// The reply would come back from the Azure OpenAI model. You'd add it to the conversation so we can maintain context.
		{Role: to.Ptr(azopenai.ChatRoleAssistant), Content: to.Ptr("Yes, customer managed keys are supported by Azure OpenAI")},

		// The user answers the question based on the latest reply.
		{Role: to.Ptr(azopenai.ChatRoleUser), Content: to.Ptr("Do other Azure AI services support this too?")},

		// from here you'd keep iterating, sending responses back from the chat completions API
	}

	resp, err := client.GetChatCompletions(context.TODO(), azopenai.ChatCompletionsOptions{
		// This is a conversation in progress.
		// NOTE: all messages count against token usage for this API.
		Messages:   messages,
		Deployment: modelDeploymentID,
	}, nil)

	if err != nil {
		//  TODO: Update the following line with your application specific error handling logic
		log.Fatalf("ERROR: %s", err)
	}

	for _, choice := range resp.Choices {
		fmt.Fprintf(os.Stderr, "Content[%d]: %s\n", *choice.Index, *choice.Message.Content)
	}

}
