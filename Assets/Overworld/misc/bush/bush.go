embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/Assets/Overworld/tilesources/overworld.atlas\"\n"
  "default_animation: \"bush\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: -0.25881904
    y: 0.0
    z: 0.0
    w: 0.9659258
  }
}
embedded_components {
  id: "collision"
  type: "collisionobject"
  data: "collision_shape: \"\"\n"
  "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"overworld\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 3.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "      w: 1.0\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 12.1965\n"
  "  data: 11.476\n"
  "  data: 3.0\n"
  "}\n"
  "linear_damping: 0.0\n"
  "angular_damping: 0.0\n"
  "locked_rotation: false\n"
  "bullet: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "shadow"
  type: "sprite"
  data: "tile_set: \"/Assets/Overworld/tilesources/overworld.atlas\"\n"
  "default_animation: \"bush-shadow\"\n"
  "material: \"/Assets/Materials/shadows.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: -12.0
    z: 7.0
  }
  rotation {
    x: -0.70710677
    y: 0.0
    z: 0.0
    w: 0.70710677
  }
}
