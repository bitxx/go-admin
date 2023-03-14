export function resolveBlob(data, name) {
  var now = new Date()
  var year = now.getFullYear() // 年
  var month = now.getMonth() + 1 // 月
  var day = now.getDate() // 日
  var hh = now.getHours() // 时
  var mm = now.getMinutes() // 分
  var ss = now.getSeconds() // 秒
  name = name + '_' + year + '_' + month + '_' + day + '_' + hh + '_' + mm + '_' + ss
  var blobUrl = window.URL.createObjectURL(new Blob([data], {
    type: 'application/vnd.ms-excel'
  }))
  const a = document.createElement('a')
  a.style.display = 'none'
  a.download = name + '.xlsx'
  a.href = blobUrl
  a.click()
}

export function resolveBlobZip(data, name) {
  var now = new Date()
  var year = now.getFullYear() // 年
  var month = now.getMonth() + 1 // 月
  var day = now.getDate() // 日
  var hh = now.getHours() // 时
  var mm = now.getMinutes() // 分
  var ss = now.getSeconds() // 秒
  name = name + '_' + year + '_' + month + '_' + day + '_' + hh + '_' + mm + '_' + ss
  var blobUrl = window.URL.createObjectURL(new Blob([data], {
    type: 'application/zip'
  }))
  const a = document.createElement('a')
  a.style.display = 'none'
  a.download = name + '.zip'
  a.href = blobUrl
  a.click()
}

/**
 * 解析blob响应内容并下载
 * @param {*} res blob响应内容
 * @param {String} mimeType MIME类型
 */
/* export function resolveBlob(res, mimeType) {
  const aLink = document.createElement('a')
  var blob = new Blob([res.data], { type: mimeType })
  // //从response的headers中获取filename, 后端response.setHeader("Content-disposition", "attachment; filename=xxxx.docx") 设置的文件名;
  var patt = new RegExp('filename=([^;]+\\.[^\\.;]+);*')
  var contentDisposition = decodeURI(res.headers['content-disposition'])
  var result = patt.exec(contentDisposition)
  var fileName = result[1]
  fileName = fileName.replace(/\"/g, '')
  aLink.href = URL.createObjectURL(blob)
  aLink.setAttribute('download', fileName) // 设置下载文件名称
  document.body.appendChild(aLink)
  aLink.click()
  document.body.appendChild(aLink)
}*/
