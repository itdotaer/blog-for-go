import { queryById } from "@/services/posts"
import { notification } from 'antd'
import pathToRegexp from "path-to-regexp"

export default {
  namespace: 'post',
  state: {
    data: null
  },
  reducers: {
    updatePost(state, { payload }) {      
      return { ...state, data: payload }
    }
  },
  subscriptions: {
     setup({ dispatch, history }) {
       history.listen(({ pathname }) => {
         const match = pathToRegexp('/posts/:id').exec(pathname)
         console.log('match', match)
         if (match) {
            dispatch({
                type: 'queryById',
                payload: {
                  id: match[1]
                }
            })
          }
        });
      },
  },
  effects: {
    *queryById({ payload = {} }, { select, call, put }) {
      console.log('query', payload)

      const resp = yield call(queryById, payload.id)

      if (!resp || !resp.success) {
        notification.error({
          description: '博客获取是吧',
          message: '网络异常',
        });
      }

      yield put({
        type: "updatePost",
        payload: resp.data
      })
    }
  },
};