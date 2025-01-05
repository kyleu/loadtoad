import {Message, Socket} from "./socket";
import {req} from "./dom";

export function socketLog(debug: boolean, parentEl: HTMLElement, terminal: boolean, url: string, extraHandlers: [...(m: Message) => void]) {
  const o = () => {
    if (debug) {
      console.log("[socket]: open");
    }
  };

  let currRow: HTMLElement | null = null;

  const newRow = () => {
    if (terminal) {
      const row = document.createElement("tr");
      const numTH = document.createElement("th");
      numTH.innerText = parentEl.children.length.toString();
      const textTD = document.createElement("td");
      row.append(numTH, textTD);
      parentEl.append(row);
      currRow = textTD;
    } else {
      const div = document.createElement("div");
      parentEl.append(div);
      currRow = div;
    }
  };

  const r = (m: Message) => {
    if (m.cmd !== "output") {
      if (extraHandlers.length === 0) {
        console.log("unknown command [" + m.cmd + "] received");
      } else {
        for (const h of extraHandlers) {
          h(m);
        }
      }
      return;
    }
    if (m.param.html === undefined) {
      console.log("no [html] key in message param: " + JSON.stringify(m, null, 2));
    }
    const html = m.param.html;

    let content = "";
    for (const x of html) {
      if (!currRow) {
        newRow();
      }
      if (x === "\n") {
        if (content === "") {
          content = "&nbsp;";
        }
        currRow!.innerHTML += content;
        content = "";
        currRow = null;
      } else {
        content += x;
      }
    }
    if (currRow) {
      currRow.innerHTML += content;
    }
    const c = req("#content");
    c.scrollTo(0, c.scrollHeight);
  };

  const e = (svc: string, err: string) => {
    console.error("socket error", svc, err);
  };

  return new Socket(false, o, r, e, url);
}
