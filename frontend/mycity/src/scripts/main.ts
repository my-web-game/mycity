import * as THREE from 'three'

let renderer:THREE.WebGLRenderer | undefined = undefined
let scene:THREE.Scene | undefined = undefined
let camera:THREE.PerspectiveCamera | undefined = undefined

function animate() {
  renderer.render( scene, camera );
}

export function init() {
    scene = new THREE.Scene()
    camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000)
    camera.position.z = 5;

    const geometry = new THREE.BoxGeometry( 1, 1, 1 );
    const material = new THREE.MeshBasicMaterial( { color: 0x00ff00 } );
    const cube = new THREE.Mesh( geometry, material );
    scene.add( cube );

    renderer = new THREE.WebGLRenderer()
    renderer.setSize(window.innerWidth, window.innerHeight)
    renderer.setAnimationLoop( animate );
    document.body.appendChild(renderer.domElement)
}

export function uninit() {
  renderer.setAnimationLoop( undefined );
  renderer.dispose();
}