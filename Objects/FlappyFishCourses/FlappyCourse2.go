components {
  id: "Coral"
  component: "/Scripts/Coral.script"
}
embedded_components {
  id: "WinCollision"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"WinTrigger\"\n"
  "mask: \"FlyingFish\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 1700.0\n"
  "      y: 360.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 10.0\n"
  "  data: 360.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite1"
  type: "sprite"
  data: "default_animation: \"coral-2\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 200.0\n"
  "  y: 1000.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/Visual Assets/SpriteAtlas/FlappyFish.atlas\"\n"
  "}\n"
  ""
  position {
    x: 700.0
    y: -50.0
  }
}
embedded_components {
  id: "sprite3"
  type: "sprite"
  data: "default_animation: \"coral-2\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 200.0\n"
  "  y: 1000.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/Visual Assets/SpriteAtlas/FlappyFish.atlas\"\n"
  "}\n"
  ""
  position {
    x: 1000.0
    y: 100.0
  }
}
embedded_components {
  id: "sprite4"
  type: "sprite"
  data: "default_animation: \"coral-2\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 200.0\n"
  "  y: 1000.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/Visual Assets/SpriteAtlas/FlappyFish.atlas\"\n"
  "}\n"
  ""
  position {
    x: 1400.0
    y: 1000.0
  }
  rotation {
    z: 1.0
    w: 6.123234E-17
  }
}
embedded_components {
  id: "sprite5"
  type: "sprite"
  data: "default_animation: \"coral-1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 200.0\n"
  "  y: 1000.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/Visual Assets/SpriteAtlas/FlappyFish.atlas\"\n"
  "}\n"
  ""
  position {
    x: 1400.0
    y: -100.0
  }
}
embedded_components {
  id: "CoralCollision"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"Coral\"\n"
  "mask: \"FlyingFish\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 320.0\n"
  "      y: 550.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 680.0\n"
  "      y: 150.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 3\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 970.0\n"
  "      y: 300.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 6\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 1430.0\n"
  "      y: 800.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 9\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 1380.0\n"
  "      y: 50.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 12\n"
  "    count: 3\n"
  "  }\n"
  "  data: 50.0\n"
  "  data: 250.0\n"
  "  data: 10.0\n"
  "  data: 62.5\n"
  "  data: 250.0\n"
  "  data: 10.0\n"
  "  data: 50.0\n"
  "  data: 250.0\n"
  "  data: 10.0\n"
  "  data: 50.0\n"
  "  data: 250.0\n"
  "  data: 10.0\n"
  "  data: 50.0\n"
  "  data: 250.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"coral-1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 200.0\n"
  "  y: 1000.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/Visual Assets/SpriteAtlas/FlappyFish.atlas\"\n"
  "}\n"
  ""
  position {
    x: 300.0
    y: 700.0
  }
  rotation {
    z: 1.0
    w: 6.123234E-17
  }
}
