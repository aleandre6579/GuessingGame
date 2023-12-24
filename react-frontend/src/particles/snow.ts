import type { ISourceOptions } from "@tsparticles/engine";

const options: ISourceOptions = {
  background: {
    image: "linear-gradient(19deg, #21D4FD 0%, #B721FF 100%)"
  },
  fpsLimit: 120,
  interactivity: {
    events: {
      onClick: {
        enable: true,
        mode: "",
      },
      onHover: {
        enable: true,
        mode: ["bubble", "repulse"],
      },
    },
    modes: {
      repulse: {
        distance: 100,
        speed: 0.08,
      },
      attract: {
        distance: 200,
        duration: 1,
        factor: 10,
        maxSpeed: 5,
        speed: .2
      },
      bubble: {
        distance: 300,
        size: 12,
        duration: 2,
        opacity: 0.8,
      },
    },
  },
  particles: {
    color: {
      value: "#ffffff",
    },
    move: {
      enable: true,
      gravity: {
        enable: true,
        maxSpeed: {min: 1, max: 3},
        acceleration: {min: 2, max: 5},
      },
      random: false,
      speed: { min: 0.2, max: 0.5 },
      straight: false,
    },
    number: {
      density: {
        enable: true,
      },
      value: 1000,
    },
    opacity: {
      value: 0.5,
    },
    shape: {
      type: "circle",
    },
    size: {
      value: { min: 3, max: 8 },
    },
  },
  detectRetina: true,
};
export default options;
