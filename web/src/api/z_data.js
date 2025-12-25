import service from '@/utils/request'
// @Summary 图表数据
// @Produce  application/json
// @Param data body {name:"string",status:"string"}
// @Router /z/List [get]
export const load_Data = () => {
  return service({
    url: '/z/List',
    method: 'get'
  })
}