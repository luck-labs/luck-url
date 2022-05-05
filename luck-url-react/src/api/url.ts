import request from "./request";

const APIS = {
    create: '/v1/api/create'
}

export function createUrl(data: any) {
    return request({
        url: APIS.create,
        method: 'POST',
        params: data
    })
}