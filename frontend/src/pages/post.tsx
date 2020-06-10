import { connect } from 'umi';
import styles from './post.less';

const Post = ({ post, loading }) => {
  console.log('post', post)
  const { data } = post

  return (
    data ? (<div className={styles.container}>
      <h2>{data.title}</h2>
      <div>PostedBy:{data.postUser}|UpdateTime:{data.updateTime}</div>

      <div>
        {data.content}
      </div>
    </div>):(<div></div>)
  );
}

export default connect(({ post, loading }) => ({
  post, loading: loading.models.post
}))(Post);
