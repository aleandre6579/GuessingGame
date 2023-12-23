import { useEffect, useMemo, useState } from "react";
import Particles, { initParticlesEngine } from "@tsparticles/react";
import {
  type Container,
  type ISourceOptions,
} from "@tsparticles/engine";
 import { loadFull } from "tsparticles"; 

function Background() {
  const [init, setInit] = useState(false);

  useEffect(() => {
    initParticlesEngine(async (engine) => {
      await loadFull(engine);
    }).then(() => {
      setInit(true);
    });
  }, []);

  const particlesLoaded = async (container?: Container): Promise<void> => {
    console.log(container);
  };

  const options: ISourceOptions = useMemo(
    () => ({
      background: {
        image: "linear-gradient(19deg, #21D4FD 0%, #B721FF 100%)"
      },
      fpsLimit: 120,
      interactivity: {
        events: {
          onClick: {
            enable: true,
            mode: "repulse",
          },
          onHover: {
            enable: true,
            mode: ["bubble", "attract"],
          },
        },
        modes: {
          repulse: {
            distance: 300,
            duration: .1,
            speed: 5,
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
            size: 15,
            duration: 2,
            opacity: 1,
          },
        },
      },
      particles: {
        color: {
          value: "#ffffff",
        },
        move: {
          enable: true,
          random: false,
          speed: 3,
          straight: false,
        },
        number: {
          density: {
            enable: true,
          },
          value: 250,
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
    }),
    [],
  );

  if (init) {
    return (
      <Particles
        id="tsparticles"
        particlesLoaded={particlesLoaded}
        options={options}
      />
    );
  }

  return <></>;
};

export default Background