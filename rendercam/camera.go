components {
  id: "script"
  component: "/rendercam/camera.script"
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
  properties {
    id: "z_offset"
    value: "100.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "nearZ"
    value: "-1000.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "farZ"
    value: "1000.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "viewDistance"
    value: "1000.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "fixedAspectRatio"
    value: "true"
    type: PROPERTY_TYPE_BOOLEAN
  }
}
