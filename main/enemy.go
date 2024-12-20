embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"yellow_character\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assetss/images/enemy.atlas\"\n"
  "}\n"
  ""
  position {
    z: 10.0
  }
}
embedded_components {
  id: "enemy"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemy\"\n"
  "mask: \"sword\"\n"
  "mask: \"spear\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 36.216854\n"
  "  data: 36.75182\n"
  "  data: 16.0\n"
  "}\n"
  ""
}
