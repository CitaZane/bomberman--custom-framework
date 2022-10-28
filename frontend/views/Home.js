/* @jsx jsx */

import jsx from "../../framework/vDom/jsx"

export function HomeView() {
    return {
        template: (
            <div>
                <label for="name">Enter your username: </label>
                <input type="text" id="name"></input>
            </div>
        )
    }
}