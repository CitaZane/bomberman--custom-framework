/* @jsx jsx */

import jsx from "../framework/vDom/jsx";

export function Overlay({ content }) {
  console.log("Content:", content);
  return {
    template: (
      <div id="overlay">
        <div class="overlay-content" id="winner">
          {String(content)}
        </div>
      </div>
    ),
  };
}
