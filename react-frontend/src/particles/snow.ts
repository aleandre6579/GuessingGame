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
        mode: "attract",
      },
      onHover: {
        enable: true,
        mode: ["bubble", "repulse"],
      },
    },
    modes: {
      repulse: {
        distance: 70,
        speed: 0.15,
      },
      attract: {
        distance: 100,
        duration: 0.2,
        factor: 5,
        maxSpeed: 20,
        speed: 5
      },
      bubble: {
        distance: 300,
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
      value: 4000,
    },
    opacity: {
      value: 0.5,
    },
    shape: {
      type: "circle",
    },
    size: {
      value: { min: 2, max: 6 },
    },
  },
  detectRetina: true,
};
export default options;
