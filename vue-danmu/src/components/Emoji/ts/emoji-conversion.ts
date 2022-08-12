import { emojiList } from "../ts/emoji-list";

export const analyzeEmoji = (content: string) => {
    //编译表情替换成图片展示出来
    let res = content;
    const allEmoji = /\[[(\u4e00-\u9fa5)|_]+\]/g;
    var matchList = content.match(allEmoji);

    if (matchList) {
        for (let i = 0; i < matchList.length; i++) {
            let matchItem = emojiList.find((item) => {
                return "[" + item.name + "]" === matchList![i];
            });

            if (matchItem && matchItem.className) {
                res = res.replace(
                    `[${matchItem.name}]`,
                    `<span class="${matchItem.className}" style="display:inline-block;vertical-align:bottom;width:72px;height:72px;zoom:40%;"></span>`
                );
            }
        }
    }
    return res;
}
