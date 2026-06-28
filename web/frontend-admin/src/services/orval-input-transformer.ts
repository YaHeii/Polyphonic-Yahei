const singleUploadPath = '/admin-api/v1/upload/upload_file';
const multiUploadPath = '/admin-api/v1/upload/multi_upload_file';

const tagSlugMap: Record<string, string> = {
  用户管理: 'user',
  分类管理: 'category',
  友链管理: 'friend',
  接口管理: 'api',
  操作日志: 'operation-log',
  文件日志: 'file-log',
  文件管理: 'upload',
  文章管理: 'article',
  游客管理: 'visitor',
  照片管理: 'photo',
  留言管理: 'message',
  登录日志: 'login-log',
  登录认证: 'auth',
  相册管理: 'album',
  网站管理: 'website',
  菜单管理: 'menu',
  角色管理: 'role',
  访问日志: 'visit-log',
  评论管理: 'comment',
  说说管理: 'talk',
  通知管理: 'notice',
};

type SwaggerParameter = {
  in?: string;
  name?: string;
  type?: string;
  items?: {
    type?: string;
  };
};

type SwaggerOperation = {
  consumes?: string[];
  parameters?: SwaggerParameter[];
  tags?: string[];
};

type SwaggerPathItem = Record<string, SwaggerOperation>;

type SwaggerSpec = {
  definitions?: Record<string, unknown>;
  paths?: Record<string, SwaggerPathItem>;
};

const strippedDefinitions = new Set([
  'UploadFileReq',
  'MultiUploadFileReq',
  'EmptyReq',
  'PingReq',
  'GetClientInfoReq',
  'SyncApiReq',
  'CleanMenuReq',
]);

function rewriteTags(operation: SwaggerOperation): SwaggerOperation {
  if (!operation.tags?.length) {
    return operation;
  }

  return {
    ...operation,
    tags: operation.tags.map((tag) => tagSlugMap[tag] ?? tag),
  };
}

function patchSingleUploadOperation(operation: SwaggerOperation): SwaggerOperation {
  const parameters = [...(operation.parameters ?? [])];
  const fileParameterIndex = parameters.findIndex(
    (parameter) => parameter.in === 'formData' && parameter.name === 'file'
  );

  if (fileParameterIndex >= 0) {
    parameters[fileParameterIndex] = {
      ...parameters[fileParameterIndex],
      type: 'file',
    };
  } else {
    parameters.unshift({
      in: 'formData',
      name: 'file',
      type: 'file',
    });
  }

  return {
    ...operation,
    consumes: ['multipart/form-data'],
    parameters,
  };
}

function patchMultiUploadOperation(_operation: SwaggerOperation): SwaggerOperation | undefined {
  // Swagger 2.0 multi-file upload still degrades to a single Blob in Orval output.
  // Keep this endpoint on a hand-written client until the upstream generation path is replaced.
  return undefined;
}

export default function transformAdminSwagger(spec: SwaggerSpec): SwaggerSpec {
  if (!spec.paths) {
    return spec;
  }

  const definitions = spec.definitions
    ? Object.fromEntries(
        Object.entries(spec.definitions).filter(([name]) => !strippedDefinitions.has(name))
      )
    : undefined;

  const paths = Object.fromEntries(
    Object.entries(spec.paths).flatMap(([path, pathItem]) => {
      const nextPathItem: SwaggerPathItem = {};

      for (const [method, rawOperation] of Object.entries(pathItem)) {
        let operation = rewriteTags(rawOperation);

        if (path === singleUploadPath) {
          operation = patchSingleUploadOperation(operation);
        }

        if (path === multiUploadPath) {
          const patchedOperation = patchMultiUploadOperation(operation);
          if (!patchedOperation) {
            continue;
          }
          operation = patchedOperation;
        }

        nextPathItem[method] = operation;
      }

      if (Object.keys(nextPathItem).length === 0) {
        return [];
      }

      return [[path, nextPathItem] as const];
    })
  );

  return {
    ...spec,
    definitions,
    paths,
  };
}
