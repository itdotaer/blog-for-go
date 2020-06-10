import api from "@/services/api"
import request from "@/utils/request"

const { host, posts } = api

export async function query(index, size) {
  return request.get(host + posts + "?index=" + index + "&size=" + size)
}

export async function queryById(id) {
  return request.get(host + posts + "/" + id)
}