import React from "react";
import ft from "../assets/endraft.png";
import ppay from "../assets/ppay.png";
import tikets from "../assets/Tikets.png";
import react from "../assets/reactjs.png";
import html from "../assets/html.png";
import css from "../assets/css.png";
import Tailwind from "../assets/tailwind.png";
import docker from "../assets/docker.png";
import js from "../assets/js.png";
import postgre from "../assets/postgre.png";
import { Link } from "react-router";

function Home() {
  return (
    <>
      <nav className=" bg-primary w-full h-20 flex gap-5 px-20 items-center sticky top-0">
        <div className="text-white font-bold hover:bg-white hover:rounded-full cursor-pointer hover:text-black py-3 px-5">About Me</div>
        <div className="text-white font-bold hover:bg-white hover:rounded-full cursor-pointer hover:text-black py-3 px-5">My Porject</div>
      </nav>
      <section className="w-full flex bg-gradient-to-r from-indigo-500 px-20 justify-center items-center">
        <div className="flex flex-col gap-7 flex-1">
          <div className="text-6xl font-bold animate-pulse">
            Front End Development
          </div>
          <div className="text-2xl font-semibold w-full max-w-4xl">
            I am a Frontend Developer with experience in building scalable,
            responsive, and user-friendly web applications. Proficient in modern
            frontend technologies such as React.js, JavaScript, and Tailwind
            CSS, I have a strong foundation in creating intuitive user
            interfaces and seamless user experiences. I am also skilled in
            integrating frontend applications with backend services and
            databases. With a passion for continuous learning and innovation, I
            am eager to contribute to dynamic teams and deliver high-quality
            solutions that meet both user and business needs.
          </div>
        </div>
        <div className="opacity-85">
          <img src={ft} alt="" />
        </div>
      </section>
      <article className="flex flex-col items-center justify-center gap-5 mt-5">
        <div className="text-2xl font-bold">My Skills</div>
        <div className="flex-shrink-0 flex h-40 items-center rounded-full bg-primary px-20">
          <img src={html} alt="" className="object-fill h-32" />
          <img src={css} alt="" className="object-fill h-40" />
          <img src={js} alt="" className="object-fill h-20" />
          <img src={react} alt="" className="object-fill h-36" />
          <img src={Tailwind} alt="" className="object-fill h-44" />
          <img src={docker} alt="" className="object-fill h-28" />
          <img src={postgre} alt="" className="object-fill h-28" />
        </div>
      </article>
      <section className="w-full flex flex-col px-20 justify-center items-center mt-5">
        <div className="text-2xl font-bold">My Project</div>
        <div className="flex items-center justify-center">
          <div className="modal-box flex gap-10 flex-col">
            <div className="text-3xl font-bold">My Film Tikets</div>
            <div>
              <img src={tikets} alt="" />
            </div>
            <div className="flex justify-between text-xl">
              <Link to="https://github.com/endradwi/fgh22-frontend">
                <div>Link Github Repository</div>
              </Link>
              |
              <Link to="https://bit.ly/MyFilmApp">
                <div>Link Deploy Project</div>
              </Link>
            </div>
          </div>
          <div className="modal-box flex gap-10 flex-col">
            <div className="text-3xl font-bold">P-pay E-wallet</div>
            <div>
              <img src={ppay} alt="" />
            </div>
            <div className="flex justify-between text-xl">
              <Link to="https://github.com/endradwi/Ppay">
                <div>Link Github Repository</div>
              </Link>
              |
              <Link to="https://p-pay.netlify.app/">
                <div>Link Deploy Project</div>
              </Link>
            </div>
          </div>
        </div>
      </section>
    </>
  );
}

export default Home;
