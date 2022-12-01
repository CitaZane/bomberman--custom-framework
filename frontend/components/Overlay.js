/* @jsx jsx */

import jsx from "../framework/vDom/jsx";

export function Overlay({ type, content }) {
  return {
    template: (
      <div id="overlay">
        <div class="overlay-content" id={type}>
          {String(content)}
        </div>
      </div>
    ),
  };
}
