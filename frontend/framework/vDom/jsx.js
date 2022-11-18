export default function jsx(type, props, ...children) {
  if (!props) props = {};
  if (!children) children = [];

  children = [].concat.apply(
    [],
    children.filter((child) => child)
  );

  if (typeof type === "function") {
    return { ...type(props) };
  } else {
    return { type, props, children };
  }
}
