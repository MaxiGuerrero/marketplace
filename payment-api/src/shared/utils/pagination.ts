const isNumber = (value: any) => !Number.isNaN(parseFloat(value)) && Number.isFinite(value);

const getNumberIfValid = (value: any) => (isNumber(value) ? parseFloat(value) : null);

const getNumberIfPositive = (value: any) => {
  const n = getNumberIfValid(value);
  return n && n >= 0 ? n : null;
};

const pageNumber = (page: number | any, size: number | any) => (page - 1) * size;

// Pagination for sequelize
const paginate = (params: any) => {
  let limit = 10;
  let offset = 0;
  if (!params) {
    return { limit, offset };
  }
  if (params.page && params.page_size) {
    offset = pageNumber(getNumberIfPositive(params.page), params.page_size);
    limit = getNumberIfPositive(params.page_size) || limit;
  }
  return { limit, offset };
};

// Sorting for sequelize
const sortBy = (params: any) => {
  let sort = 'updatedAt';
  let dir = 'ASC';
  if (params.sort) {
    ({ sort } = params);
  }
  if (params.sort_dir) {
    dir = params.sort_dir;
  }
  return { sort, dir };
};

// Pagination for mongoose
const pagination = (params: any) => {
  let limit = 10;
  let page = 1;
  if (!params) {
    return { limit, page };
  }
  if (params.page && params.page_size) {
    page = getNumberIfPositive(params.page) || page;
    limit = getNumberIfPositive(params.page_size) || limit;
  }
  return { limit, page };
};

// Sorting for mongoose
const sorting = (params: any) => {
  let sort = 'updatedAt';
  let dir = -1;
  if (params.sort) {
    ({ sort } = params);
  }
  if (params.sort_dir) {
    dir = params.sort_dir === 'ASC' ? 1 : -1;
  }
  return { sort, dir };
};

export { pagination, sorting, paginate, sortBy };
