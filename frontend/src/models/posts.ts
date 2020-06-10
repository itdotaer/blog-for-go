import { query } from "@/services/posts"
import { notification } from 'antd'

export default {
  namespace: 'posts',
  state: {
    list: []
  },
  reducers: {
    updatePosts(state, { payload }) {
      return { ...state, list: payload }
    }
  },
  subscriptions: {
     setup({ dispatch, history }) {
       history.listen(({ pathname }) => {
          if (pathname === '/posts') {
            dispatch({
                type: 'query'
            })
          }
        });
      },
  },
  effects: {
    *query({ payload = {} }, { select, call, put }) {
      console.log('query')

      const resp = yield call(query, 1, 10)

      if (!resp || !resp.success) {
        notification.error({
          description: '博客获取是吧',
          message: '网络异常',
        });
      }

      yield put({
        type: "updatePosts",
        payload: resp.data
      })
    }
  },
};