import { toast } from "react-hot-toast";

// const API_URL = "/api/v1/";
const API_URL = "http://127.0.0.1:21101/git-contribution-draw/v1/contributions/";

// export function fetchData(username) {
//   return fetch(API_URL + username).then((res) => res.json());
// }

export async function fetchData(username) {
  const res = await fetch(API_URL + username);

  // if (!res.ok) {
  //   throw new Error("Network response was not ok");
  // }

  const data = await res.json();

  // 后端返回的结构化错误处理
  if (data.code && data.message) {
    const errMsg = `{"code": ${data.code.code}, "message": "${data.message}"}`;
    const error = new Error(errMsg);
    error.statusCode = data.code.http_status_code;
    error.type = data.code.name;
    throw error;
  }

  return data;
}

export function download(canvas) {
  try {
    const dataUrl = canvas.toDataURL();
    const a = document.createElement("a");
    document.body.insertAdjacentElement("beforeend", a);
    a.download = "contributions.png";
    a.href = dataUrl;
    a.click();
    document.body.removeChild(a);
  } catch (err) {
    console.error(err);
  }
}

export function downloadJSON(data) {
  try {
    const dataString = JSON.stringify(data);
    const dataUrl =
      "data:text/json;charset=utf-8," + encodeURIComponent(dataString);
    const a = document.createElement("a");
    document.body.insertAdjacentElement("beforeend", a);
    a.download = "contributions.json";
    a.href = dataUrl;
    a.click();
    document.body.removeChild(a);
  } catch (err) {
    console.error(err);
  }
}

export async function share(canvas) {
  try {
    canvas.toBlob(async (blob) => {
      navigator
        .share({
          title: "GitHub Contributions",
          text: "Check out my #GitHubContributions history over time. A free tool by @sallar and friends. https://github-contributions.vercel.app",
          files: [
            new File([blob], "contributions.png", {
              type: blob.type
            })
          ]
        })
        .catch(() => {
          // do nothing
        });
    }, "image/png");
  } catch (err) {
    console.error(err);
  }
}

export async function copyToClipboard(canvas) {
  if ("ClipboardItem" in window) {
    // https://bugs.webkit.org/show_bug.cgi?id=222262
    // https://web.dev/async-clipboard/
    const item = new ClipboardItem({
      "image/png": new Promise((resolve) => {
        canvas.toBlob(resolve, "image/png");
      })
    });
    navigator.clipboard
      .write([item])
      .then(() => toast("🎉 Copied image!"))
      .catch((err) => {
        toast("Sorry, copying image is not supported on this browser");
        console.error("failed to copy");
      });
  } else {
    toast("Sorry, copying image is not supported on this browser");
    console.error("failed to copy");
  }
}

export function cleanUsername(username) {
  return username.replace(/^(http|https):\/\/(?!www\.)github\.com\//, "");
}
