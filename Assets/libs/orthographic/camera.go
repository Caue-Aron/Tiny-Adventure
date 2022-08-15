components {
  id: "camera"
  component: "/Assets/libs/orthographic/camera.script"
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
    id: "center_x"
    value: "96.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "center_y"
    value: "80.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "center_cam"
    value: "false"
    type: PROPERTY_TYPE_BOOLEAN
  }
  properties {
    id: "near_z"
    value: "-10.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "far_z"
    value: "10.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "zoom"
    value: "4.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "projection"
    value: "FIXED_ZOOM"
    type: PROPERTY_TYPE_HASH
  }
}
