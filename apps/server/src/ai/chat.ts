import { createOpenAICompatible } from "@ai-sdk/openai-compatible"
import { DEEPSEEK_API_KEY } from "@/config"
import { stepCountIs, streamText, tool } from "ai"
import { z } from "zod"

const deepseek = createOpenAICompatible({
  name: "deepseek",
  baseURL: "https://api.deepseek.com/v1",
  apiKey: DEEPSEEK_API_KEY,
})

const deepseekV4Flash = deepseek("deepseek-v4-flash")

const result = streamText({
  model: deepseekV4Flash,

  prompt: "读取 package.json，然后告诉我项目用了什么框架",

  tools: {
    readFile: tool({
      description: "读取项目中的文件内容",
      inputSchema: z.object({
        path: z.string().describe("相对于项目根目录的文件路径"),
      }),

      execute: async ({ path }) => {
        console.log("读取文件:", path)

        return Bun.file(path).text()
      },
    }),
  },

  stopWhen: stepCountIs(5),
})

for await (const part of result.stream) {
  switch (part.type) {
    case "text-delta":
      process.stdout.write(part.text)
      break

    case "tool-call":
      console.log("\nTool call:", part.toolName, part.input)
      break

    case "tool-result":
      console.log("\nTool result:", part.output)
      break
  }
}
