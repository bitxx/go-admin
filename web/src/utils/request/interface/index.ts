export interface Result {
  code: number;
  msg: string;
}

export interface ResultData<T = any> extends Result {
  data: T;
}

export interface ReqPage {
  current?: number;
  pageSize?: number;
}

export interface ResPage<T> {
  list: T[];
  count: number;
  extend: T;
  pageIndex: number;
  pageSize: number;
}
