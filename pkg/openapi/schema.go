// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	externalRef0 "github.com/unikorn-cloud/core/pkg/openapi"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+y9aXPiup44/FVcPFN1ZqogbdZAv5khbIGEfefSDyXbAgyy5Fg2W1d/939JssEGQ0i6",
	"77196varczrIWn767Zu+R1RimARDbNPI1+8RE1jAgDa0+L8WFnFMXWt5f2R/0yBVLd20dYIjXyN5ycH6",
	"mwMlPlSqFh8i0YjOfjGBvYxEIxgYMPLVmykSjVjwzdEtqEW+2pYDoxGqLqEB2Mz23mRDqW3peBH58SMa",
	"IcCxl4mWRTYatKrFW/vAkhgsmRbZ6Bq0ru/FG1EtfnQ71gJg/QDYmu/uxjf2+laCM35wO6ZFVlC1b+9E",
	"ckfdBIeY5kPL/xCDIbWfiKZDji6qBYENK+yqO+I3/leCbYj5/wLTRLrKT/tlRdkGv/uW+C8LziNfI//f",
	"lxNKfhG/0i8cf4aWbkOxdvCUT0TbS97eJZtIYicSEGj5cHGyH1F3sy1x+Pu3C3fAMBFk/2tAG2jA5nt3",
	"YWnsYy44Iz+iEWpC9UhGNPL1HxFN02AmB+cxRUknYqlsPB3LwRyMAahl0xk1I6cf55Fv7ID3gcVd7Cpg",
	"3ONJ7lVJJ/IOhQkiCx2/D4tdbLvdxubEMmKOhSBWicYmCQAHGkBHka+RFYEPCiKLBf0/oBrwQSVGJBqh",
	"NrAZvOC+tlQqqt7Ua9X+oRpv6FVaxZ20WqhmqmtzNCjUcg9wXztow6re1Ku7+qouN3rjZLO43lb1ra4Y",
	"ZXvS5YM3oJJadCo5xP4OhmW5uiK7Rq+UqK/q6Xqxup+3H7pz9LLbdmrdOnx5KSfavdR8a9ZhbZ7MtJrr",
	"zL42mAGtTek2rUbuvgQ/1Joc8jTsLqqYQUwwBAxVSCmw9gxZLUgJ2jBsnUMNWsCGmtTtNo+cLPSqTqyR",
	"DfnV5Bac/eN0d8GNQ89gkzX8ReimIh1ie6ZrDKnSWS2Vk2Esk5hnY6kcSMaUR02OKTkFKpl4WgOKEolG",
	"2DRBFCy35U71tT/oMQwaJzvp6oroXaT12b8nw/SK/bvdq8Yba63Y61Zp1Rhswb6agfuapT2vxRx79vfG",
	"XtOrmSrK241edce+hxyly7oqp5f9+NN+nBynO4MaHRplq/k8KKqJgdxLlBOgV0sp3bgNRuXWcDXYtI1y",
	"o5MwbVVOFxRdToFSNtXu54pKpZNoDupJrYj2Wu+ppBSXQDmUS2pvuWuW6ulh35SHldocyGP9tVDjZ2kP",
	"+8lBN15U1zYdJzu15mh8qMsd2huWaVeePE3WubFaiLfhIHeYyON0b6UBIKcb7XWn2FkPXhS5bHX28XIP",
	"L3vqoZqol9IGNBapLq7hLn7qKP1yefi83ExkkwyfzcR4OKm3u7Xca6FmgWGbU/DkeZlUE7mXPpqU2sau",
	"NzZ2m66RY+eo9da1rVap9ZREfNRHTxN1nX6Fw0a5Pch1GAy1Z7Q93gmWHx4cq2Mou+fETMHZ1zoCD+Ot",
	"DJJv1H6u51/wDmzX1TG2n9VNs7ACu9VhM4jXkDGuxxKFnlKI64mBnaeN6gtponItnXlONOSsWR/nmuYk",
	"oTrrwnMr/tTe0Zc6VVPxwRZVJ+PNqmwdhtUSLJJyLlE2zEKnMjzYzlZdPg21x1apPTbnsFauJZ7gAqiV",
	"JWy/zTujUTLdaRT3sUlTTWnDtbMpW4Nstevks7HHmQofn0Ei3bU6TrcDrN68Pnt6zcedYn7WyuWHqyXd",
	"V16aL4ny2gHFvjwyRuh1WDxktBftZZ/r1OzODPf7KkUrG1SN2mjVaLTyRu0tLuNaWo6XXmbVTD33lOx1",
	"+tYbQM0nI7Wmj7GNUZ4t1FKcguYmkVf1Uq6VeKqv1UwyvQbFZCH9jPbDXi7dXWuZwqy8Nc1Vu78Z98fy",
	"/rH0lmiYeDBfj1JOt2Vk5/1iSrG6q8oQP9cbpewhVU/MWqieeulO8jp87Rj1/Gqc3g2zo/HMKYysNFZi",
	"2a6Rn7ViaFUYNFut/Kg4Ku1AYtfdKfnaxhq/DaFTSVQ3+XVBBkrGJCv01jfWneGmOUrbeNQGm/SmmXhr",
	"5heFcX/ZrQ5HBzk2zi7VQ6ffXRR7+7aRzu37j7u3wVtB328Ly8UINZOJl+1yia35666BrPpTKj1qosOy",
	"1oqryWJh8TgZPirNWfsxL2crq4012vWMx0W/aMVWVBvmlr2u3qi1ndns0K2XW4NBo/eGD/F6sVyFDtUz",
	"lZqeGxTk/Iw4I6ot1cYLzqxgtTjIabi+K6grpd1Lv9FC6Y3E+mqhsnmWZ9sUKCxNpNUX2edKC/a7kyV4",
	"6r7G95jOqnIhl88XyzCnGaNGZlt4fnKytcI+1kuVCRx10KD7MnAqiUpNz9L5IV8uLzP6y7I92j0b6ZdG",
	"fqYT66k2KDW7o6T2mnlp9kdzjT7Ne4dFEtRJaW8mlFquAYBqV4zyvjap52Cmvutm+7tFI/PyDB8rmqPK",
	"jUp5/2Q5yQKqvyWeDuqyuVMOxfaM6Okx6Tq7V3NRQcmdXps3cAG9lXtvo3rtMe101/KsuX5ZbIxnCHLt",
	"SgcAukuP8q9dE5gzdV2YbBrjVWVGJsuUnIq99FYmSOi1RamhHmC/lyinVm/pnFUo5PvlyWC+d5Jv9lMe",
	"1gyYGiyWWOltQLVXU8wyfOrvu4vxi+pU2g/Opl1f6aivZ2uqtq/A5KsC7IXL9GcbaOlznanLkcmwLdcr",
	"tdWkMt43esv1pDje1xPtbePQ3jd7Y7lRqcuT4WRVP/TTk1XHqBfXh8lqsG4Ua+vGarBsrPK7SXF8mPQG",
	"6/FhLNeNxmrSJpFoZGEBbM9ctZkJQ2K5iv6MSx4mDzXdgqo9cyw98jWytG2Tfv3yxZVqTFNy5fAXFSCk",
	"AHV9v07iF603dJJmnktpPtpTEqOSSjB1kC3ZSyhZEMENwLbkDgVYk5rVYkFiyq0+d2U0lebEkuaOZS+h",
	"JWnQBjoKVzEdU/tNbASxk5s2ghjS9Blpv1zR8s398a0HzcwbR/hj5pxgwi1XahJMhdUKVBWaNtQ67h8v",
	"jegepwMx9RJQSYEQS95nnCK2OkKSAqW5g+Y6QuyvdI/VpUUwcSjaP0zxmDiSAfaSSRByKYsSx1Ihn8Ag",
	"WLeJJek2lZhx5AiKYrBAUFzuj2gEqMi/y/vv0I8m/CesmUTnPp9/fA/4aGgkGiEms0I4z/j6D9dQ5nAE",
	"WsTDqAi7AwRtGPn241s0whVvoCkJmIqnYnEA57FUOpuLKblkMgY0ORNPagntMTuPnBwXfO3Qneh4bgFq",
	"W45qOxa8tiPfwtlMBgI5HUtk0ulYKq6osWw8mY5puZySSUAtpcBM5Me3u/ENqCgMzfIS0qktkTm/fEoZ",
	"o7Qtwjgduxw/j79yTTbc2V+WtoEiX7+Hzs9MSXbzhjAMBc65ZiEzHBkhSnOLGByFHMosK7b2BugIKEgw",
	"VvoZJPnH94imUxOBfUNcQl4zdKxT2wI2seiZK49+IcaeqM+kYkFzyYB7N2yDWw2H89GJJ9yKOubHDTK7",
	"H9GIAjSXpX2OLKBlEYvj2wYgXZu5AI9ExS+z4La861AYO3Y/uV8ii7VCTtvxTzsHOuMe4iOJL8F3H5WI",
	"5XINMVojkEqY2BwHgY6nGBz5iiso5jpEmsBNleA50tWfBJM3yxX4gBNL2+r2km+GAgNKDHEkgBj/2Etw",
	"p1Ob/hK4uYt526JiWYAJ00KikkMdgNBespc6lQwIMGVb2ktLsIHBzXEYzYml6JoG8c8B6TjNFSgxmpVU",
	"C2oQ2zpAVNIIv8fjro73Z1r6RkdwAekvxLItoJIGsQ41SdlLLs+iLo4JSIE9E2UqcKgYxDYVGDjFQmF0",
	"t63jRXDjVCUm5PILYCnfqh6Rl5+dYS7+63TgKT55w05Hloig+SMnMBGwGWPkd7UQ+uNn7smvznBBohPc",
	"0zlnS8iJVExOx5LxXjz1NR7/KsuTiCdi4qnEPJOKxx7nWiaWAkouBrKP6VgylVIfoQI0NZ088UgQsyEw",
	"ImchAzaPlsrIspaBMZjLpGMpJZWKgaycjWVTcyUxB8nMo5yIREWMguoE63jR5fqAF7hgf4SaX++yCILV",
	"olC87pn9W5QjIf9gDlT4f2K33DX7AaXMvQOghbPw25x78TNy6s8VvnOF3z50h/S2riPGcHIOvUqmzpQY",
	"2/mkzrOAGFq6Kj336q8S51+SCRaCJ+vYhhYGqAutDbRurXInd6Z8opn4ZziDdkWwTVyzWEVAN34BB85j",
	"ycFwZ0KVWQ3ioERVHcuCWpD1gsBI2wKY6hDb7jcAa1PMRlJHVSHUGKdkote29g9SdS5m0jmLZfelAgqj",
	"kokgoIxFm8SyJd2WAGXL6JQ6AtKr7Zp+DrRruBfKu2ptmC0WSyfikWhkzWkuru22lNQ6g+IT6iqI1MjW",
	"zlUbT6atdIkx7LTGVuNlr5byszb7xt5HvkZKhQhHb3Zd+iISjewY9CrDvOK8PGEsv43oKqtr2nA5WaVj",
	"k149VU5paasGXxQFNSsDNZbGtUa/Q1vK4zpWX5berFw7r6dXL1h7RGtj/dxPGBigLW23XiLRCFszn4dm",
	"AQ272Tp5fS0c3urthIKSL9tD+RF2x69LtWvRdXY9djqg0UilDTxw2vQ5lWw3q6+lp/RoBJ6X+263sxgU",
	"gFHfTob9bd7axNcfsToYbIdQeYH7LrTDabLWbTakLVSkNdxLFNoPUk/ctgTYPxm5MvalSaajIF1lw5jS",
	"A2wJWOz259CCWBVinc01xWwyjueUzQV9H0oqwAwbuRpgE4l7zfbubC5tMG2C6gvsKQo6nWJGYroqsAoT",
	"u0wcrP0c0WJiz+ZsmisU69OLoXZSQo8qMlc5fwEF9zGzWxgs5jrWpJO+y896HpH7bZWTBSELBGM610Dt",
	"/b9IxIkIXbXImEI8LmdyyWQ2k0nFTKLKajauLejc0SzZUhxzJTvYsVbqxo4n4AMwTfog9sxkngtMN5DM",
	"GZjl854CVSUOtr0v+LD7dZrzO7yi3ISENy8Q4LdWbv4DUeDb53DgHeXoDA+EtU1MiHWtQPBcXzjWLX/Q",
	"TziOL9cIdfCbEFeL3EQ+jhMcnZtnjk1imk5VsoHWnp3n6I7jXknqmExfgNoUA7Qglm4vDfHLHALbsaB7",
	"3oCD/Pc1ylQDxlRimR9EW40YgOdNuQksfAL/mXsiyOOOi/oTzL5GlAxUE6m4FoNJJRVLgbkaA48AxNKy",
	"PFeVTFrWsvAjHCoA6+v86VxL9//h97a7fudb+vaZa3qPhfiHPkhSnVCbe0qoRJfEQUzBR0zbgxLBMMqI",
	"EloS0AwdU6ajMY0LSgCh4ExHB4rfb2J6caDflkh9oaV/jVA6Ra3gPC3nVC0eUx9hJpZKzLMxAOBjDMjx",
	"eCalqRlZUz8RtbruInEH+G/mtybNv8HdfPvg5bxDm94ofkUWQfBfdz9nEU6XigXdH6MyD5+7SI8nxAQX",
	"uetW7gcth1M4XF0blY/g3vpr/kG6pzY0mn806v84jVrcvKdRa+Eatpve8hkpJkK3Mz7Dn/TPP+mff9I/",
	"/6R//kn//A9J/4Q7U7cgnTFDMZmRZSbqQ0VB/9Df1XVRfLLUyjkyHjUI4z1apfbcQOVnuE4PJ6X0XF1N",
	"MmO5dOig8r59QKhhDFpK32w1ksjqrsq0V37aNfo1ucPlRTk+KVQzw301Pe6pu+awv5t048txbxF/7XWW",
	"9VXJHveq+3pXPtRXHdQ4LJKT4WTdOCz0UZfJoPgSDLdsg29KYum8Gp3NpP+ElGHZVArplZKQGa9H8Dmv",
	"N1elRLNXijcO9VTjUKJVAy21QjVT743T9V471Ti0k/XuVgejxoGdCzx3ZPW5nnnd5yxtWEOqkUZaZXB4",
	"NQaHcWKJVKNBleRg/Wo0Ngo7C34yx8lOXDX6bD9Ee+5s1QPZvCa1pLZPY9UoJ8ajzlLV+b4249FkqVXK",
	"+9fD0mgY/XRjVU02KvX9eFgzGqtSctyrp5tFDTUOHdQc9pONnsYLedTkQOf7M3JE0dNrJTHIu3Bwxomc",
	"zeRAfrzrkvx27bzMn0wzTeLUNPL7t8Ny3e08ZpbKqhxvFl5gSn/tZp4Krdy+OxnDQWz9VNBkO6lqmcFO",
	"aabLg3at1bGza/ktm7XURLyW7+0H2XVXbWArFl+VjXzNGTUzCyAn4i+9ThtXMtli9jBp5F63Rr3bWSaf",
	"W2W7+ZZ6LahGu9RNAA3W9pRUcrmsYdhOb2um5nlrC3gu3dyCdPnvUkN66vOT3CnJyUmiM1BLtUEjQRKd",
	"ZAf31ul9pxRf142cOXkm8cawcajrcUstmR0g73qdfu2p25v0NNROd1EnA4vaqC6v9/1+rqSt00XluVzX",
	"Kstm41lLdktL0C8OSoN4uQQM+aSG9HNWW06v1fVg2InX9MGhnG6WtZfOarntJ5/qwGi8jVe1VGNYOoz7",
	"y3azhFKjw+RplGwc+om43CwNDmPUqSvFck9ddcZdmY1L7QcJE4PBONGpmINuRauN5TgZ4lq6v487jYJf",
	"DakdOvFxCsjV/XjdmQ8O+dRkUKuqq9qok+i06pXlbmCkR/2+XQalTm8wzMW10TjZKaUtvxqiDdMmSOT2",
	"ih5fKZVcfFJIb1RD3ai4bQGsyVxFaVYfs6OsKi/3XdWaFR8fMpWF/ZrqqjUri1JkRx77G7COvYxIw7b7",
	"xfbOmODqWq0Vs20TzGCtuc10V8PnZKGbW6H1pFNYJLXHfvzRjiky3cTi8aFjDFF/89gp08eUUgJrK9eH",
	"iVh3oC2cIsi/Ppe03KKweW29DTJPRvs12bVIebgYOI91qMt9WScWzJRi8CU2U+xHo9KX5cao0tssWvX1",
	"uDJZb61RFqq17B6sXmNxOxZrxPeLXqeShMV+Cq8bpVqpnIrbb0+5ZWFM6SzfNwq4SuVOGZgDJ/a4fFms",
	"Mr2D1sSZ/La1shyw325QdXdYlc16dQiUBennW4c3MOs2LVSJgcduLl53kstD51FJo3Ir0ctWOinSIUva",
	"b1idiZ2rLiZOvlZRB/XHlCHbqeRkU+u+FDtpGRqPsUPNSqdTbxoCo+ybk1jaO3vcf0LFWOuw26bo1jG2",
	"sWQyXa8dAB21KoWS1SvOU/DQHT0VlCpNV59TqtKZtQ7205uyHvQmiXHL2T+qzU71pa0fsgjVJ4WtbtEE",
	"0B6fnzcOei0v6ijd7WfQJnNY6rH2uKfIWm+jZovqy/Oyglb7YtsujPe7UjlWcfrJwUgvPmdx5bmGjMQg",
	"3VmBjtEz2+tVHs8ST7k+yj5lt9tuvNNsFrTewFRVrQviZTmlH6ppOO4149UU3dlA2easWElOZPcZbdC0",
	"jW7LVOdglc2WnnKzsdZKwuzQWmj9gzyrtUpE2w/7HQOnq5gUKhnSHG8cMh/o3VEtNWraq3rpcbNc4NS+",
	"PW8iqPSwMkCDzGGcGSAl8dTCj4PRoFfIbw5V25hv0LicVBepmLOOx9ex116325YNDaFMZoG33ee3VaNd",
	"NdZ4vTUHhZ5hOCZEq4qstId9O15L0FSzscGvuFXOWghjqzl8Kmy2uJ5Mas3Ecp/b2jLUzJdYtZ5ElW5L",
	"T+qjeKrUSxGzjPWJ8jpRerpZ2LYmh00XVpaoDkej3mGRfnMa7YZjbu2qVl6MjRpQcVKOww7pPDS75lv+",
	"sao56/xj7PnVrqcKnX474hqTXqXGEwQWtD5YaBFq0Tr2EmLbtURFpNzhtufcQTyGYkHbsTDPwwik14nY",
	"ugjHe5nqIuOD8Ml5UqOOVeRoPFeEl2N4jgc3MK/PRaBepOKxxY8Bcm5AO9hLV4Y/GZx37WmRU3gtMzQI",
	"C5Fh8ytSasLm9ZIN3RxtAY8loJJQ/tzzU2jpeE4+6UNwNJ93xAP9g4P1NbFwTEXE0R6I5Wqcka/xZDye",
	"yMZzj0zhBLb3B1n8gdJ7J6OOcqWk+m5Yegd/r0IZKMSxT0mbx5RzP1QffHErt5QChbUgCM+aF7gpwokm",
	"tAydUh4BES4wE1q221ZggYgC0B05+6VjLcFZwOuOb7tsM1pghqOL+MNfv/J8kx9Rr2sCUbxaGd8uwyDl",
	"+kfPAOWDjgjbnnJc2AKXEBM+yfAyFvdDPiR63tYhWGnx7rGbp8FeQwhRaPOPiDu/b7pvt8FBP4w53Plt",
	"Q4N+ADUip0sBlgX27iaOBwnt5XG2+vFMbAMQO8ZdVTIhsA5C8P7jn4D6ERCczhgOg3MKuNjPMeGeg187",
	"5grcpl7on/AjBKxr76BwtShyHYWM4ywbMs70cInXZ8jJG9GcNnYFM8OI+sYlWXeC5977CuFIF7fm2Ms6",
	"tJckBFRdL3dDOpO+Bv+ABpBXdE+gULWgPTMJL4EJ/lEBVFcZYBGdub+wecPxOlDiE0pSVa0lEosv8CVQ",
	"hHQJ7aVjACwxIuO5eCLLXNRwhOzkOiNUASZYVwG6PUUYVwtFl7Oqpht4cjz6B7AhCNAQTBBJyjeQgA8I",
	"XDpTYqJunxb3v15ltOYqLtHIHBg62s9cbr7QNxB7/2CKjNBeohFEVICgFw2LRkzdq9xjKksYjqhEg4Ul",
	"QAjiBXwfh9lwSfXGh+FwN5HOsJUR0HHoiq7Cer5GxU0/F8nVBqQULGCUV14CW2c4xquKRBgqhL2Fz5qX",
	"bGhR6M4qNsEUUIA19n9ubsZzr9dyh7ADPkg8w53yZF0FUJHbzQa6xdeBmuuopDgir1fMCzWxU7Y/S4c2",
	"sPZeKSmbXKgO+VaVSrw+itkHbHJCoTevyOIVa/lBe1ke5zcfXHYQiV6YAg4+JpDNPENG2FnR45ycWTI8",
	"CSbm29AwiQUshnwOPuK/78Pjqt4feMH92aq+IvxoII3YV8omcGnGfgUIke3F1g2o6cCb5FTedRXHgsbP",
	"OWYMoKUwmLuYJolfFa+Iis/wPiO6XsAQxp04GET+03US44OOCuWJS4R0LnAFgK+K7cKlGQYdzvV4mcv3",
	"29UmUrUYZI7hxUWipdnFMue88VSodHWm7RIKoXJJ4P7sgXszvoTo9vQltnTdm8UXvH+3MKfLBp7f/HE/",
	"7kTh9+19fe3El+f0AuwnAfZO+sxp9KubXn+sYrr9rbgp7yO3jun+T87g4S16FQ6ia8DtqxfaO3cdCMX9",
	"JxDB0wP5uv+Ge3+fvO5XPHw1fpd0dapXuVzRX6jyIHUhDDYTqQ1fupJGVMeA2HadN+ENRK6wgMD8kRBQ",
	"XPwhWF1zc0JfZQ2vt2dzSewfcK67pS4AS3AnatOkpKVJJrBsLm6xBiyNTrFKDEO3bQgfpEJYO5W7Dh9E",
	"QFFo9f2+m/NdzsXVhYFn7SjQwtCG9BUoEA0AckKJhsta6cVRIB8sITaa/dWBUQYoplKj/SnFPOBwcGuR",
	"RGHRFOtYgzuoeUWiDNZM6+EnBzaDbuRr5P//hxzL5WMTEDt8++///Xr6V2z28O27HM3Ef/hG/M///leY",
	"SAhr0XdxuFfefMErsBe+1jCDVnQ2DDMqGC/7i0p8hAQ0zYKUT4AdJNQXt63mxf48dho+LdyZSNXtUy20",
	"J6Tfndftsxg2Kf/J00yZsoP0i3Jvt32pZ8m+r5aI9cJYU0hlTZhd6Cq6x5P+doI5eI5Pc+qQae7oaXsJ",
	"hFMCYKiRy391YwMnu0Hkky/JVjAeX0I5E4CBzrVki3knAjf5boo5Ae+Jww0PgN3P5sR6mOIwyhNb6HIn",
	"ws1NCj/Dv3ZzXqJj2LZEVIUP8DO27RJavATQhJZIyHVsIp1qaCyIgK1vYLgfyvvDR3CM6+0XHiyx8+gJ",
	"Ad7HsnADgJ2VCzgyPwWR/AjnWQIi95MZTLpqEUrmdqiSH9Y69G5K/230sH8Kmd/UzC7zXO/U0cJqFi+V",
	"tbB6sYvdhFWLXV5A0CSEV8MZDLPcUx19oZxCdazbOmDaQ9A1+SDsawPYka8Rx9LD+QnQDTo7Wue3QErP",
	"vF/3QvRY8X/hZSManB29UDPXC3XfZsI9WKc93n3fYZ6zUKvXc318GFr801PHBxFB9SsBd9oOntMhZHde",
	"2t2M6gus48UMoMWMa5H37lZ8J/nqEk8HYDuvFr269ns37E6Z92YM3fctiSGop8qHSP+tUypaMP1PqCBY",
	"bddUdH68m27gziQUUq9G3wMBL/Lnl8Uv6tQ06h56CvjlPowp3tdcggS0e19BvodH15XKu67HW+saSomg",
	"4EdP4AL6gxFF4bUM2YRA6s/wRHpstAB36hIwHuEmhvhiN5yJHK+a3nPBwR3xoM3HGBd9J5D0s3zCF8QK",
	"AaiXr/AeTHVNPUHUg+QCMmFzntQAcCCL4S/q9tF7D5bXVLAr8vACGcIO4+MDIRgcIu5uUOy9d31NNtzL",
	"le+Qg6G60Hnl8rvPZvwC48/6GVPPbzh+WgO8z/pkahVCzTkvWPvEgb6fwee8+CyUZvyw9mxEaAWbkyoQ",
	"Eby4L6x+tuglOL6dA+Sq3XvW1vYsROyWeof6SrjvRYxg3OuyutBH5ktC7WuJMMLcKTjUJoYoEj9fbsjQ",
	"k0Lb7bjkAZCKtkvuh0wP8RikmFIyAAYLbuT4tjXF3N8IeYsZGky1ivIcKYcyGy0qVcQsblsuwFv/iLti",
	"x7UcezllLE/qPOULoZrHZVH8/XTgCV9/8XzYNVSLHvBPTivCjhAVnRd5ygZ3qdrQMjyvqkNdgHvWzMMU",
	"V+fSHCDxoU5FDyMGFAlIiqMjxuCPa0R53qPliFUwHxVEpin2YqgSweE5A95knFrvjYOIwbcIgkPuPQ4R",
	"bqIXRYjQQwvHXvqFMGbMNxTVJSmvLYl6Jren2OuFCbw+AyKIrCKH6huI9hLQNKhJGx1w16Ou6rYbLDGg",
	"oUCLLnVTkvJYkwyH2lPMMBNIfzGhEdMx++4v1336IElFQYoXusPZFgywF91vp5itay+hbgUdqVG+mucs",
	"VgEnB5fOLj0XNMpnn2IDmFSkL7pMQSDFpatDkgIUeLZh1x08xdRRl4zqdCMIGWCaOmOUftI7BlHZNUSi",
	"HuP69g5RXneaXIrH38dj8qul5W1vSaBpxt2ukovmKZcKp79rw43ntH47x7S7sU+D3v/99XOHZOf6gvrv",
	"hhLZuNDwU6DD+3tg/21w/hfB/Camn1pQ3InkgbYjofh9Frm/vbivD/ftnIzg2LsyMy5E7TW5tyRbxrcR",
	"IWvHDKgVUYESIkc8KglRdFQMvN6i3ugphg+LB1eF+lL3HNnuJH6CZFOdKw9eZmeod1wkqZ/lnX8L97hw",
	"jPPr/lcSbT0ccqOuPPcbISnfqp4UdMaDRBL9lqFwSMDyZjpSIKPW95OrbvkqToCzYMqGUEJ4S2+ebmYQ",
	"i7fgtuHOvpmDeQtnQ6PP4cmXPgi2QjplhJ3P31HDjXuSeaB94+keHbzGZIvP+nD4/8ktYA2e/SwSsm7d",
	"9+fsPbZdXe2c48ylxcfzzI+dTUKjPLoBg8adaAaOoMtBj3YR46kxNvyqfnwB9Xt4bMh9nd9xyOQhZmT0",
	"gwTDaeTBfxlBdv+HAj9OgSdf7I2cwqB/OJByLjIJRbKgrwxfOJekwD98P7rVbWGj/X/FBIeXXFgEwWuq",
	"HfvtV3uarikFYdqA6Jt0q8KADbjfX+4d9Zqv/NbFnfzh3oWJAKIvY9y0yFxH4WC+CKXcWuoikuNfttRN",
	"xxPha4Rzxg970m7wVFfRfNpfT/eRtktyVEj9zDWUEoPtp+5n0u4C9zPpa8Uz7mvNvhqa0H0aROP54O+e",
	"3NW/3z+5N+M7JwfBc/vU+3vOHVbmEwD5HbKkJ14Lcdm+TgNZbK7LYuVQt6u0cM1pBP9le+9fTDHA+6DO",
	"wcYsIUD20s3IF7n7CsRwrtvCOwG87EH9ALUpPu5AnDuQwHJiGb682FsRHD7qpvaOJWApum2JogE2/C4F",
	"3m1fcZFHAALPylEHhRiOwUZcoQ5qEzBkDZachmXd+7q6XMUsHUsUqgRr1Bes4u92IUBtniR0mlvHNlyI",
	"MvBTg5iwJJZqsXCMND+Ex1cDjT5uHNMdeWMuf6H6+UTPzEI6r3rWuaNUYRowpBDb7rMGfu4i6i7e9+37",
	"1o4Gry4A/zC5FvYW4cX+iR9l/DHEsHy3WRh7K4g8smrxQep4L9Ec3yGackVhGjmrcng3gVI8Evw9pO78",
	"FMkWpTu/dk3fG5VXzskVH2/Ur1ze/3DlXdUiklRn/FCBEtR5BvNxZcL+3wSUbomlTSPh7nb35xtKONli",
	"aEnewPCznlb56HmDz3BegbY3SOp3qr8S2O+wh3yQKxwjSYHPPromE97hhZFnEPcG/lqInzEWH7qFMY9j",
	"v4JL2AQ4HbYtQk2oHl/JC8tbu6IUMW7P6+fEXMDRdIiv6DG8n0OokBGci7HCPZc4V8QJsK9EEn1SCVCR",
	"eapdmYPSW3sQiQnhSUe2fuvLK9ogVua3vqI2sOwbR+ZdK8I+P8vCoI7iuZjfyTIP1JR6uMJWgqpj6fa+",
	"y1R8NxDO5UqwUUho4qPbG99dh3r5qQrvB3PEjIt0RUS2l0mRBVdqBP7Yt9BdfT68R383iS+B748JIpGv",
	"3z2p+ok5vYY1J/jxn0Q/kHBKY3fVF1NKVS9ulm9VPUcoPXah4e+V6owZ8Jeb5kAV6UoOdX0VAKEp9uZy",
	"X4pxW4BYZKdD+iBJeSrp9l/0RJbsa1cRNxxk6zEbYrYGP94Ua9BEZC88I7otAdWm7vNGYLGw4EJcLAJ7",
	"aPE5RPnrqWpHPGfs7iU6xZpOTWCrS6a5I39NCD1pTy475J8qQF1DLGhVtxnfi4RBKxKNbKBFBUjlh/iD",
	"7OXlAlOPfI0kH+SHpCh9WXKU+vKwhQjFuD/Sfe8hpt5O4K0aJoICEnxrx7x0trlFWAZ+hzuvbC/h/ZTI",
	"boIF9JIa9sI+P3tf8thAISrepPIligVM+WO7Ct4FuQLtIUTohZ2qGZKUfPY6b0KWrzk6juO+3Hpw4wdH",
	"7C/A1L9s4i4dmDcSshGjaDdHg0HS15PYdf9fZGpLvFXTMaLsotIUM/X7WDhgwYVObXgUomExaG5SH4nK",
	"oRJwqyOm2B0rMnQ4vRBKdQXt+ctPqgXdPMItlDAUer6PW0Hfc7FXEIEZkdQ76WU4PfQi86Y+iDfPQPqZ",
	"C7zd7vtHNJKS4+/PEtrUin+cfP/jy+c/f0Qj6Xs2f+uROr9U4h6pcHn0j28/vgXQ9J04eBiSnr8j0Qy8",
	"BgEseEpRcksYp9ifLeWicSCBypeVahHHhiH8EEwxY4AxiE+BMckXZWNGwDHQ5q5xHGeAvcSL7qc4+ESN",
	"Wycg2Y6FxQh69ogNmUtzHcPYwgI8fUgoE3we3iTn2D/Be4jylAZ6bGYlABLIAcnbnrE8xUD4udyEZp5e",
	"w/ViBmwMz8KFwE2H4Q/4uYkyEpm7RUkuiK9zYbcVXHiCw+kJN17y9MXLIHOr9Ty6vkKfAVz6FHsNfTvm",
	"Z6nyt6CtL98DSYvFH3+I7V9GbNeIoQLt88j7ndjdDNzlT+P6f5oAYrqnBQxoi5dlw5c9DfkSJJ2W9wN3",
	"qpuOfZWUztMqhI/S5+sPXnXLue+qOZk8EW1/HWLeEJ1dFl+vGbhv0d/mRzjinJmMgcSQY2dPtD/GRE49",
	"Pt2mk//JSswFo/0S2iyyAvnTc1K+8CoBSomqc1Ae3U5M1EU9Rdy1xsLyqH281CvU5R3MeGqmxO18g9gw",
	"aILaxGuxN8X90fvC+qxL49HZdHTTfUAwB/E5r6JPcS+gor+nfP51rOcDCOh1lYpd65vSIghRn/Dym2rs",
	"ji+k3QcUPDd//6jZec1Fwt7ZPnbIEIqCX4J+HLXOOtV9Cs2Cc/yRk/98ZL3a2uc97dRL3ZTEXXnI4/c/",
	"8CR5U8cLLxIt8NrX7iacBGzC7CAXh7c6QlPMjB9NO/kfTlEyUbHo+rW916I9/ssMr2NMUjwh7SqsDoUh",
	"SqVNeP49Py8TBryn3qkFJ/0n0eGnCO4n6Gzxh7x+Vg0l9Loe6qWGiGxQDLen1OUzDZTQj1zzBxVRkdxT",
	"Ea22rmmg8Tux5TdAlpSce/9LrwHj30BRFVT45Tv/r679OGbbwquYxYW1m5DLsQuLdCHGp67gWJHPeB+W",
	"VcROIvfYKXxs0EA5pvz+O3Ek9f6XmNhl4uB/p1b6R9j+hLD1zH2J6vhYsuplJBGEwoz9uwXqbRr4e3DK",
	"v6dYjb77qcspP+ARCuDGz3iCLpHjU26hd6TxfXz2d3ME/T2Y7v2C+XPB3LCKrvM+e3+riO4U+0kAoJDo",
	"rfDq89Z4jMR4z8c9cSyJbHEg/4EHE6a4ubaBeIBoEOeZHqIrnwjw+zj6STp5UklaAqwhsTu2CwtOseu2",
	"8MEL2Jc+Dn4V74SnzzzHf9GPRaqvs41fEcQmv134+u+iZf2z7b0Ct7DuwZ7P2X5h2PNBsXPexe8n7MDz",
	"qf4YhP9cufPl+6kbyk3jUJh4n0PD+8zDM0Rsnbq03KPCVC+k1x+z8V9jNv5RQ34/NeRfZi2d+DUn1Tus",
	"pj43LD4pzxz757jIP0mwvRNiD8Onk6H4hyP9Ktl2vSWJL7fY68bi42heIP1BarkzRCVAJdWxLIhtLj24",
	"g2uKXQ/Xy7Hg3Ps2KvKm3L5KjNADYf5gtcAUnwdmrpoNPHPZO9evi+t45/yUreBt548T6p+j61M3oONh",
	"6rHe0Z+a8Rd9Bwvu0f0DaPCpyE/La9hzjTUm7sanIEbcca8K0Nx1/9gJd6PwLhZ4vUWkj0S+Rnhzt4+z",
	"Wm4/sP+7y3ygXvsjjtdhOVJ+9UzzEF3w0im+8tK0SJkCkgqoCsQTdceSKN7ZhOlUmmu/ICQyP0UuqIoc",
	"pk65wsASqVTuH+lnDRmPqloeZCKfoQomMkz7HD3/fmTx9zB6riXzMrzytY/7tKC9gQvyfwyH/I9y6L1v",
	"Ox05591W001svMM4CsfGTwWZ3hX8ITaR+01ooOnhD1L/DibUtT5OYV4gPla4LYSB9M7b77zZruVLV0M6",
	"Xh9tJIP37AWnFJObOXC+td0SJ54y7dpnP20fdTgYPsOq+cb+WEY/k1R6rYb8ZsWwqHgPaojiLZiruNQ6",
	"VoAH2yUzHTA4EaAMl6++Zuwl7pvAsnXVQcDingB4eiPxWJkPTt02+MNV3l2IGrzWS6H0MMVj4vCUFrHI",
	"3u1ZIQqFpxGvATmWiOUm1yzBhnvU3FdcCgRjqJ78o/tjNyMRx5c0h7sveRsW78WOUDoR3rRBIn9WzP+J",
	"tGv/DH5ETQpl+LJ3iNu0RAHq2nOjuM/DMMUdIMR/6Xde7xcdS9tAZ1h+J6WEfClqs0/Iutqu6ftvcwb6",
	"A91O88US71XFe9mKV3P4azlebQjvn7OXasPeRcrUFB9zpqAV5X4DuAMMJU+OSGJJIfYUvYkGNXbEz9w+",
	"g80v5k5n0Of1hiGINLfhyWkiLXWb+t/zocH3fKISkDQdILJgFqW/8dQUL6Dtd76ctXf3RB7v8u9SmR9z",
	"Txl0XpRAdIpnhquYI0DQXt1PoMbS8Bo78I4FzKBdcE+6Y/uT50SHeWJx4gn1qYf7nrrQay/GpfRfxwRA",
	"LXjW8xMdM/NCPVAe5vD3ST+jbfpfP73QMt/jHe5WT0/EX7rh+53qwyU2XWlpdClwLlhoOHADouaEd24D",
	"0GDjs9MTSudS5/QAhNtE/9TQ6PKpc6+10YNUnfsetBKsX/hSTjJPiJ+g+JDOpAd3owBEiXghy00JkyQm",
	"/K72VOKeGfZN8IEzYJrIXZq6z1I6VOyLM0vFIltGr27t0JmAniOylbbEQRrbim6YFlDZjyjA6qZYhBQd",
	"mxhCZhDDYMdETFV1X+4Uxb82IUjHi6i0JFu44TAXTiVM7Cm2IPtStHMBvE2L95iZ7xF5D9fyraoAJia2",
	"CPmJXUi25bALmOJT85Y7Mzg8Guq5Les+TEP+JnbXLLXbrNid4d/vfvj1guN6ky5PEAf7c737Jte7gv0T",
	"812VyH1v95+5Ve/ov+h2PmiQEl1Tv3h63E1ee+zwc2zJLB6kdL+9CvBeMFTXdSWvy7rPVHSNQE7vvrz3",
	"vSTkyDm3vNCdHiSpaks6pjYEmuRJZvelnyO9+/Txy4db3NffgJdTfOISlyJjiu0AE/Z4T8hZGSfyBIrL",
	"evEZYw/HL11TC97dfFDq+vUej5F72R2Xh3n4ZWT948f/CwAA//8YtGcwEtIAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "https://raw.githubusercontent.com/unikorn-cloud/core/main/pkg/openapi/common.spec.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}