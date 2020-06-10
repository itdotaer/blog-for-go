import { connect } from 'umi';
import styles from './posts.less';
import PostItem from '@/components/PostItem'

const Posts = ({ posts }) => {
  return (
    <div className={styles.container}>
      {
        posts && posts.list ? (
          posts.list.map(post => {
            return (<PostItem {...post} />)
          })
        ) : (<div>123</div>)
      }
    </div>
  );
}

export default connect(({ posts }) => ({
  posts,
}))(Posts);
