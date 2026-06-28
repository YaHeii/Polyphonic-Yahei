import { defineConfig } from 'orval';

export default defineConfig({
  adminApi: {
    // 1. 输入源：直接读取后端生成的 swagger.json
    input: {
      target: '../../service/api/admin/internal/docs/swagger.json',
      override: {
        transformer: './src/services/orval-input-transformer.ts',
      },
    },
    // 2. 输出目标：严格限制在 generated 目录下
    output: {
      clean: true,
      mode: 'tags-split',
      target: 'src/services/generated/api.ts',
      schemas: 'src/services/generated/model',
      client: 'axios-functions',
      mock: false,
      // 3. Mutator 桥接层：将生成的代码与业务 request.ts 对接
      override: {
        mutator: {
          path: 'src/services/client.ts',
          name: 'customClient',
        },
      },
    },
  },
});
