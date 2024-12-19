import request from "@/utils/request";
import { ReqPage, ResPage } from "@/utils/request/interface";
import { UploadRequestOption } from "rc-upload/lib/interface";

export interface FilemgrAppModel {
  id?: number;
  version?: string;
  platform?: string;
  appType?: string;
  localAddress?: string;
  downloadType?: string;
  downloadUrl?: string;
  remark?: string;
  status?: string;
  createBy?: number;
  createdAt?: Date;
  updateBy?: number;
  updatedAt?: Date;
}

export const getFilemgrAppPageApi = (params: ReqPage) => {
  return request.get<ResPage<FilemgrAppModel>>(`/admin-api/v1/plugins/filemgr/filemgr-app`, {
    ...params,
    pageIndex: params?.current
  });
};

export const getFilemgrAppApi = (id: number) => {
  return request.get<FilemgrAppModel>(`/admin-api/v1/plugins/filemgr/filemgr-app/` + id);
};

export const addFilemgrAppApi = (data: object) => {
  return request.post<object>(`/admin-api/v1/plugins/filemgr/filemgr-app`, data);
};

export const updateFilemgrAppApi = (id: number, data: object) => {
  return request.put<object>("/admin-api/v1/plugins/filemgr/filemgr-app/" + id, data);
};

export const delFilemgrAppApi = (params: number[]) => {
  return request.delete<object>(`/admin-api/v1/plugins/filemgr/filemgr-app`, { ids: params });
};

export const exportFilemgrAppApi = (query: object) => {
  return request.download(`/admin-api/v1/plugins/filemgr/filemgr-app/export`, query);
};

export const exportUploadFileAppApi = (options: UploadRequestOption) => {
  const { file, onSuccess, onError } = options;

  const formData = new FormData();
  formData.append("file", file);

  return request
    .post<object>(`/admin-api/v1/plugins/filemgr/filemgr-app/upload`, formData, {
      headers: {
        "Content-Type": "multipart/form-data"
      }
    })
    .then(response => {
      if (onSuccess) {
        onSuccess(response.data, file);
      }
    })
    .catch(error => {
      if (onError) {
        onError(error);
      }
    });
};
