import axios from 'axios'
import { getToken } from '@/utils/auth'

const mimeMap = {
  xlsx: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
  zip: 'application/zip'
}

const baseUrl = process.env.VUE_APP_BASE_API
export function downLoadZip(str, filename) {
  var url = baseUrl + str
  axios({
    method: 'get',
    url: url,
    responseType: 'blob',
    headers: { 'Authorization': 'Bearer ' + getToken() }
  }).then(res => {
    resolveBlobWithHeader(res, mimeMap.zip)
  })
}

export function downLoadFile(str) {
  var url = baseUrl + str
  const aLink = document.createElement('a')
  aLink.href = url
  document.body.appendChild(aLink)
  aLink.click()
  document.body.appendChild(aLink)
}
/**
 * 解析blob响应内容并下载 ,
 * @param {*} res blob响应内容 其中包含网络请求的header
 * @param {String} mimeType MIME类型
 */
export function resolveBlobWithHeader(res, mimeType, fName) {
  const aLink = document.createElement('a')
  const blob = new Blob([res.data], { type: mimeType })
  // //从response的headers中获取filename, 后端response.setHeader("Content-disposition", "attachment; filename=xxxx.docx") 设置的文件名;
  const patt = new RegExp('filename=([^;]+\\.[^\\.;]+);*')
  const contentDisposition = decodeURI(res.headers['content-disposition'])
  const result = patt.exec(contentDisposition)

  let fileName = fName
  if (result != null && result.length >= 1) {
    fileName = result[1]
  }

  fileName = fileName.replace(/\"/g, '')
  aLink.href = URL.createObjectURL(blob)
  aLink.setAttribute('download', fileName) // 设置下载文件名称
  document.body.appendChild(aLink)
  aLink.click()
  document.body.appendChild(aLink)
}

export function resolveBlobNoHeader(bobFile, fName) {
  const aLink = document.createElement('a')

  aLink.href = URL.createObjectURL(bobFile)
  console.log(aLink.href)
  aLink.setAttribute('download', fName) // 设置下载文件名称
  document.body.appendChild(aLink)
  aLink.click()
  document.body.appendChild(aLink)
}
