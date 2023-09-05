# citi

Citi bank CSV transaction exports don't include cardholder name, but we can get it from the web UI.

Go to the transactions view and select the statement period you want to
download. Then you can use some scripts in the devtools interpreter to load all
the transactions and their details, and then finally download the html.

# scripts

Load all transactions:

TODO: does it really need multiple loads? only mobile maybe?

```javascript
f = document.querySelector(".footer-message:not(.mobile):not(.ng-star-inserted)")

while (f.textContent == "Load More Transactions") {
    f.click();
    f = document.querySelector(".footer-message:not(.mobile):not(.ng-star-inserted)");
}
```

```javascript
do {
	f = document.querySelector(".footer-message:not(.mobile):not(.ng-star-inserted)")
    f.click();
} while (f.textContent == "Load More Transactions");
```


```javascript
while (true) {
    f = document.querySelector(".footer-message:not(.mobile):not(.ng-star-inserted)");
	if f.textContent != "Load More Transactions" {
		break
	}

    f.click();
}
```

Expand all transactions details:

```javascript
document.querySelectorAll(".transaction-row.desktop.ng-star-inserted > .transaction-description > .cell-wrapper").forEach(tx => tx.click())
```
