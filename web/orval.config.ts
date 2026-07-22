import { defineConfig } from "orval";

export default defineConfig({
    api: {
        input: {
            target: "http://localhost:8080/openapi.json",
        },
        output: {
            target: "src/api",
            schemas: "src/api/model",
            client: "react-query",
            mode: "tags",
            clean: true,
        }
    },
});