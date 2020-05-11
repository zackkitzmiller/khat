import React, { Component } from "react"
import "./UsernameInput.scss"

class UsernameInput extends Component {
    render() {
        return (
            <div className="UsernameInput">
                <input onKeyDown={this.props.updateUsername} placeholder="Username" />
            </div>
        )
    }
}

export default UsernameInput
