const refs = new Object();
const onMountedStack = new Array();

function renderElem(node) {
  if (typeof node === "string") {
    return document.createTextNode(node);
  }
  const $el = document.createElement(node.type);

  for (const key in node.props) {
    if (/^on[A-Z]/.test(key)) {
      $el.addEventListener(key.slice(2).toLowerCase(), node.props[key]);
    } else if (key === "checked") {
      $el.checked = node.props[key];
    } else {
      $el.setAttribute(key, node.props[key]);
    }
  }

  // append all children)
  for (const child of node.children) {
    $el.appendChild(render(child))
  }

  if (node.props.ref !== undefined) {
    refs[node.props.ref] = $el;
  }

  return $el;
}

const render = (vNode) => {
  // console.log("vNode", vNode)
  // return text node if element is just a string
  if (typeof vNode === "string") {
    return document.createTextNode(vNode);
  }

  // on initial render vNode can have template and onmounted properties
  if (vNode?.template) {
    if (vNode?.onMounted) {
      onMountedStack.push(vNode.onMounted)
    }
    return renderElem(vNode.template);
  } else {
    return renderElem(vNode);
  }
};

export default render;
export { refs, onMountedStack };
