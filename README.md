# 和暦(wareki)

## 使い方(Usage)

```go
locationJST := time.FixedZone("Asia/Tokyo", 9*60*60)
t := time.Date(2023, 1, 1, 0, 0, 0, 0, locationJST)
w := wareki.Get(t) // 令和5年
```

## 取り扱う元号一覧

- 明治: 1868年10月23日 ~ 1912年7月30日 [1]
- 大正: 1912年7月30日 ~ 1926年12月25日 [2]
- 昭和: 1926年12月25日 ~ 1989年1月7日 [3]
- 平成: 1989年1月8日 ~ 2019年4月30日 [4]
- 令和: 2019年5月1日 ~　[5]

[1] https://crd.ndl.go.jp/reference/modules/d3ndlcrdentry/index.php?page=ref_view&id=1000255208

[2] https://www.archives.go.jp/ayumi/kobetsu/m45_1912_01.html

[3] https://www.archives.go.jp/ayumi/kobetsu/t15_1926_01.html

[4] https://www.archives.go.jp/naj_news/16/anohi.html

[5] https://houseikyoku.sangiin.go.jp/column/column104.htm
