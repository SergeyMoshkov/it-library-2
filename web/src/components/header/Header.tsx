import logo from "../../assets/react.svg"
export default function Header() {
  return (
    <div>
      <div>
        <h1>
          <a href="/">IT Library</a>
        </h1>
      </div>
      <div>
        <p>IT Library 2</p>
      </div>
      <div>
      <img src={logo} alt="logo"/>
      </div>
      
    </div>
  );
}
