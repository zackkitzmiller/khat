import React, { Component } from 'react';
import { connect, sendMsg } from './api'
import ChatHistory from './components/ChatHistory/ChatHistory'
import ChatInput from './components/ChatInput/ChatInput'
import Header from './components/Header/Header'
import UsernameInput from './components/UsernameInput/UsernameInput'

class App extends Component {

  constructor(props) {
    super(props)
    this.state = {
      chatHistory: [],
      author: "Anonymous Coward",
    }
    this.send = this.send.bind(this)
    this.updateUsername = this.updateUsername.bind(this)
  }

  componentDidMount() {
    connect((msg) => {
      console.log("New Message")
      this.setState(prevState => ({
        chatHistory: [...this.state.chatHistory, msg]
      }))
      console.log(this.state);
    });
  }

  send(event) {
    if (event.keyCode === 13) {
      sendMsg(event.target.value, this.state.author)
      event.target.value = ""
    }
  }

  updateUsername(event) {
    if (event.keyCode === 13) {
      this.setState({author: event.target.value})
      event.target.disabled = true
    }
  }

  render() {
    return (
      <div className="App">
        <Header />
        <UsernameInput updateUsername={this.updateUsername} />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChatInput send={this.send} />
      </div>
    )
  }
}

export default App;
