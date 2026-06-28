const invalidUploadPaths = new Set([
  '/admin-api/v1/upload/upload_file',
  '/admin-api/v1/upload/multi_upload_file',
]);

type SwaggerParameter = {
  in?: string;
  name?: string;
  type?: string;
};

type SwaggerOperation = {
  consumes?: string[];
  parameters?: SwaggerParameter[];
};

type SwaggerSpec = {
  paths?: Record<string, Record<string, SwaggerOperation>>;
};

function isBrokenFileUploadOperation(operation: SwaggerOperation | undefined): boolean {
  if (!operation) {
    return false;
  }

  const parameters = operation.parameters ?? [];
  return parameters.some((parameter) => {
    if (parameter.in !== 'formData') {
      return false;
    }

    return !parameter.type || parameter.type === 'array';
  });
}

export default function transformAdminSwagger(spec: SwaggerSpec): SwaggerSpec {
  if (!spec.paths) {
    return spec;
  }

  const paths = { ...spec.paths };

  for (const path of invalidUploadPaths) {
    const operations = paths[path];
    if (!operations) {
      continue;
    }

    const brokenMethods = Object.entries(operations).filter(([, operation]) =>
      isBrokenFileUploadOperation(operation)
    );

    if (brokenMethods.length === 0) {
      continue;
    }

    delete paths[path];
  }

  return {
    ...spec,
    paths,
  };
}
