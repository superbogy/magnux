import {Provider} from '@tarojs/redux'
import Taro, {Component} from '@tarojs/taro'
import '@tarojs/async-await'
import action from './utils/action'
import Index from './pages/index'
import dva from './dva'
import models from './model'

import './app.scss'


const dvaApp = dva.createApp({
  initialState: {},
  models: models,
  onError(e, dispatch) {
    dispatch(action("sys/error", e));
  },
});
const store = dvaApp.getStore();

class App extends Component {

  config = {
    pages: [
      'pages/index/index',
      'pages/collect/collect',
      'pages/search/search',
      'pages/tags/tags',
      'pages/article/article'
    ],
    window: {
      backgroundTextStyle: 'light',
      navigationBarBackgroundColor: '#0068C4',
      navigationBarTitleText: 'ogre magi',
      navigationBarTextStyle: 'white',
      enablePullDownRefresh: true
    },
    tabBar: {
      color: "#626567",
      selectedColor: "#2A8CE5",
      backgroundColor: "#FBFBFB",
      borderStyle: "white",
      list: [
        {
          pagePath: "pages/index/index",
          // text: "首页",
          iconPath: "./asset/images/berries.png",
          selectedIconPath: "./asset/images/berries_focus.png"
        },
        {
          key: "new",
          pagePath: "pages/index/index",
          iconPath: "./asset/images/camera.png",
          selectedIconPath: "./asset/images/camera_focus.png"
        },
        {
          pagePath: "pages/collect/collect",
          text: "收藏",
          iconPath: "./asset/images/herb-1.png",
          selectedIconPath: "./asset/images/herb-2.png"
        }]
    }
  }

  componentDidMount() {
    dvaApp.dispatch({type: 'sys/test'})
  }

  componentDidShow() {
  }

  componentDidHide() {
  }

  componentCatchError() {
  }

  render() {
    return (
      <Provider store={store}>
        <Index />
      </Provider>);
  }
}

Taro.render(<App />, document.getElementById('app'))