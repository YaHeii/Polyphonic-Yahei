import assert from 'node:assert/strict';
import fs from 'node:fs';
import path from 'node:path';
import { pathToFileURL } from 'node:url';

const projectRoot = path.resolve(import.meta.dirname, '..');
const transformerPath = path.join(projectRoot, 'src/services/orval-input-transformer.ts');
const swaggerPath = path.join(projectRoot, '../../service/api/admin/internal/docs/swagger.json');

const { default: transformAdminSwagger } = await import(pathToFileURL(transformerPath).href);
const rawSpec = JSON.parse(fs.readFileSync(swaggerPath, 'utf8'));

assert.ok(rawSpec.paths['/admin-api/v1/upload/upload_file'], 'raw swagger should include upload_file');
assert.ok(
  rawSpec.paths['/admin-api/v1/upload/multi_upload_file'],
  'raw swagger should include multi_upload_file'
);
assert.equal(
  rawSpec.paths['/admin-api/v1/upload/upload_file'].post.parameters[0].type,
  undefined,
  'raw single upload param should be missing swagger type'
);
assert.equal(
  rawSpec.paths['/admin-api/v1/upload/multi_upload_file'].post.parameters[0].type,
  'array',
  'raw multi upload param should keep explicit swagger type'
);

const transformedSpec = transformAdminSwagger(rawSpec);

assert.equal(
  transformedSpec.paths['/admin-api/v1/upload/upload_file'],
  undefined,
  'transformer should remove invalid single upload path'
);
assert.equal(
  transformedSpec.paths['/admin-api/v1/upload/multi_upload_file'],
  undefined,
  'transformer should remove multi upload path that generates invalid TS client'
);
assert.ok(
  transformedSpec.paths['/admin-api/v1/account/find_account_list'],
  'transformer should keep normal admin paths'
);

console.log('orval input transformer test passed');
