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
  rawSpec.paths['/admin-api/v1/upload/upload_file'].post.parameters[0].name,
  'file_path',
  'raw single upload should currently only expose file_path from goctl swagger'
);

const transformedSpec = transformAdminSwagger(rawSpec);

assert.deepEqual(
  transformedSpec.paths['/admin-api/v1/upload/upload_file'].post.consumes,
  ['multipart/form-data'],
  'transformer should rewrite single upload consumes to multipart/form-data'
);
assert.equal(
  transformedSpec.paths['/admin-api/v1/upload/upload_file'].post.parameters[0].type,
  'file',
  'transformer should rewrite single upload param to swagger file type'
);
assert.equal(
  transformedSpec.paths['/admin-api/v1/upload/upload_file'].post.parameters[0].name,
  'file',
  'transformer should restore the missing single upload file field'
);
assert.deepEqual(
  transformedSpec.paths['/admin-api/v1/upload/upload_file'].post.tags,
  ['upload'],
  'transformer should rewrite upload tag to ASCII slug'
);

assert.ok(
  transformedSpec.paths['/admin-api/v1/upload/multi_upload_file'] === undefined ||
    transformedSpec.paths['/admin-api/v1/upload/multi_upload_file'].post.parameters[0].type === 'file',
  'transformer should either keep a compatible multi upload file param or drop the path explicitly'
);
assert.ok(
  transformedSpec.paths['/admin-api/v1/account/find_account_list'],
  'transformer should keep normal admin paths'
);
assert.deepEqual(
  transformedSpec.paths['/admin-api/v1/account/find_account_list'].post.tags,
  ['user'],
  'transformer should rewrite normal tags to ASCII slugs'
);
assert.equal(
  transformedSpec.definitions?.UploadFileReq,
  undefined,
  'transformer should drop upload request definition leaked from interface{}'
);
assert.equal(
  transformedSpec.definitions?.MultiUploadFileReq,
  undefined,
  'transformer should drop multi upload request definition leaked from interface{}'
);
assert.equal(
  transformedSpec.definitions?.EmptyReq,
  undefined,
  'transformer should drop unified empty request definitions from frontend input spec'
);
assert.equal(
  transformedSpec.definitions?.PingReq,
  undefined,
  'transformer should drop legacy empty ping request definition after EmptyReq unification'
);
assert.equal(
  transformedSpec.definitions?.GetClientInfoReq,
  undefined,
  'transformer should drop legacy empty get client info request definition after EmptyReq unification'
);
assert.equal(
  transformedSpec.definitions?.SyncApiReq,
  undefined,
  'transformer should drop legacy empty sync api request definition after EmptyReq unification'
);
assert.equal(
  transformedSpec.definitions?.CleanMenuReq,
  undefined,
  'transformer should drop legacy empty clean menu request definition after EmptyReq unification'
);

console.log('orval input transformer test passed');
