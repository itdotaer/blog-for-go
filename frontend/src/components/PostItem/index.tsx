import React, { forwardRef } from 'react'
import { Card, Tag } from 'antd'
import { Link } from 'umi'
import styles from "./index.less"
import { hydrate } from 'react-dom';

export default class Post extends React.Component {
    render() {
      const { id, title, postUser, description, content, createTime, updateTime } = this.props

      const titleNode = (
        <div>
          <h2>{title}</h2>
          Posted By : {postUser}
          <br/>
          <Link to={`posts/${id}`}>ViewDetail</Link>
        </div>
      )

      return (
        <div>
          <Card title={titleNode}>{description}</Card>
        </div>
      )
    }
}