import Taro, { Component } from '@tarojs/taro'
import { View, Text } from '@tarojs/components'
import {connect} from '@tarojs/redux'
import { AtTabBar } from 'taro-ui'
import './index.scss'

@connect(({feeds, loading}) => ({
  ...feeds,
  isLoad: loading.effects["feeds/getNews"],
  isLoadMore: loading.effects["feeds/getMoreNews"],
}))
export default class Index extends Component {

  // constructor() {
  //   super(...arguments);
  // }
  config = {
    navigationBarTitleText: '首页'
  }

  componentWillMount () { }

  componentDidMount () { }

  componentWillUnmount () { }

  componentDidShow () { }

  componentDidHide () { }

  handleClick() {
    console.log('---->')
  }

  render () {
    return (
      <View className='index'>
        <Text>Hello world!</Text>
        <AtTabBar
          tabList={[
            { title: '待办事项', iconType: 'bullet-list', text: 'new' },
            { title: '拍照', iconType: 'camera' },
            { title: '文件夹', iconType: 'folder', text: '100', max: '99' }
          ]}
          onClick={this.handleClick.bind(this)}
        />
      </View>
    )
  }
}

