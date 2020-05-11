import React, { Component } from "react"
import "./Message.scss"

class Message extends Component {
    constructor(props) {
        super(props)
        let temp = JSON.parse(this.props.message)
        this.state = {
            message: temp
        }
    }

    render() {
        var msg
        if (this.state.message.author === "") {
            msg = <div className="system-message">{this.state.message.body}</div>
        } else {
            msg = <div><strong>{this.state.message.author}: </strong>
                {this.state.message.body}</div>
        }
        return <div className="Message">
                {msg}
            </div>
    }
}

export default Message
