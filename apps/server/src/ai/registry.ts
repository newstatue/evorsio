import { createOpenAICompatible } from "@ai-sdk/openai-compatible"
import { createProviderRegistry } from "ai"

import { DEEPSEEK_API_KEY } from "@/config"

const deepseek = createOpenAICompatible({
  name: "deepseek",
  baseURL: "https://api.deepseek.com/v1",
  apiKey: DEEPSEEK_API_KEY,
})

export const registry = createProviderRegistry({
  deepseek,
})
