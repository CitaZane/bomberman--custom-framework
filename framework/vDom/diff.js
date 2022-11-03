/* -------------------------------------------------------------------------- */
/*                                    diff                                    */
/* -------------------------------------------------------------------------- */
// calculate the differences between two virtual trees
// return a patch fn that takes the real DOM and perform appropriate ops to patch it up

import render from "./render";
import { onMountedStack } from "./render";
import {refs} from './render'

const zip = (xs, ys) => {
  const zipped = [];
  for (let i = 0; i < Math.min(xs.length, ys.length); i++) {
    zipped.push([xs[i], ys[i]]);
  }
  return zipped;
};

const diffAttrs = (oldAttrs, newAttrs) => {
  const patches = []; //store all patches
  //remove attrs if not present any more
  for (const [k, v] of Object.entries(oldAttrs)) {
    // remove old listeners because the attributes are not changing
    if (/^on[A-Z]/.test(k)) {
      patches.push(($node) => {
        $node.removeEventListener(k.slice(2).toLowerCase(), v);
        return $node;
      });
      continue;
    }
    if (!(k in newAttrs)) {
      patches.push(($node) => {
        $node.removeAttribute(k);
        return $node;
      });
    }
  }
  // setting new Attrs
  for (const [k, v] of Object.entries(newAttrs)) {
    patches.push(($node) => {
      if (/^on[A-Z]/.test(k)) {
        //add new listeners
        $node.addEventListener(k.slice(2).toLowerCase(), v);
      } else if (k == "checked") {
        //add checked attribute
        $node.checked = v;
      } else if(k == 'ref'){
        refs[v] = $node;
      }else{
        $node.setAttribute(k, v);
      }
      return $node;
    });
  }
  return ($node) => {
    for (const patch of patches) {
      patch($node);
    }
    return $node;
  };
};

const diffChildren = (oldVChildren, newVChildren) => {
  const childPatches = [];
  // recursevly call diff for children to get all patches
  oldVChildren.forEach((oldVChild, i) => {
    childPatches.push(diff(oldVChild, newVChildren[i]));
  });

  const additionalPatches = [];

  // take the (additional children) from new tree elem
  // append children
  for (const additionalVChild of newVChildren.slice(oldVChildren.length)) {
    additionalPatches.push(($node) => {
      $node.appendChild(render(additionalVChild));
      return $node;
    });
  }

  return ($parent) => {
    // since childPatches are expecting the $child, not $parent,
    // we cannot just loop through them and call patch($parent)
    for (const [patch, $child] of zip(childPatches, $parent.childNodes)) {
      patch($child);
    }

    for (const patch of additionalPatches) {
      patch($parent);
    }
    return $parent;
  };
};

const diff = (oldVTree, newVTree) => {
  // CASE : newVTree is undefined
  //  simply remove the $node passing into the patch
  if (!newVTree) {
    return ($node) => {
      $node.remove();
      // the patch should return the new root node.
      // since there is none in this case,
      // we will just return undefined.
      return undefined;
    };
  }

  // in case of component, unwrap the template as newVtree
  if (newVTree.template) {
    // save component onMounted function
    if (newVTree.onMounted) {
        onMountedStack.push(newVTree.onMounted);
    }
    newVTree = newVTree.template;
  }

  if (oldVTree.template) {
    oldVTree = oldVTree.template;
  }
  // CASE: They are both TextNode (string)
  // CASE : One of the tree is TextNode, the other one is ElementNode
  if (typeof oldVTree === "string" || typeof newVTree === "string") {
    // If they are not the same, replace $node with render(newVTree).
    if (oldVTree !== newVTree) {
      // could be 2 cases:
      // 1. both trees are string and they have different values
      // 2. one of the trees is text node and
      //    the other one is elem node
      // Either case, we will just render(newVTree)!
      return ($node) => {
        const $newNode = render(newVTree);
        $node.replaceWith($newNode);
        return $newNode;
      };
    } else {
      // If they are the same string, then do nothing.
      return ($node) => $node;
    }
  }

  // CASE : oldVTree.tag !== newVTree.tag
  // assume that in this case, the old and new trees are totally different.
  //  instead of trying to find the differences between two trees,
  //  just replace the $node with render(newVTree)
  if (oldVTree.type !== newVTree.type) {
    return ($node) => {
      const $newNode = render(newVTree);
      $node.replaceWith($newNode);
      return $newNode;
    };
  }
  //  Now oldVTree and newVTree are both virtual elements.
  //  They have the same tagName
  //  They might have different attrs and children
  //  patch attributes and children
  const patchAttrs = diffAttrs(oldVTree.props, newVTree.props);
  const patchChildren = diffChildren(oldVTree.children, newVTree.children);

  return ($node) => {
    patchAttrs($node);
    patchChildren($node);
    return $node;
  };
};

export default diff;
