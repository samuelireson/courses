---
book: abstract-algebra
chapterName: Groups
chapterNumber: 1
---

---
An introduction to the fundamental topic of groups. We look at some examples, and prove some results.

---
## Contents


## Basic axioms and examples
> **Definition**: 
> With respect to a set $ G $, a **binary operation** is a well defined function,
> $$
> \begin{align*}
> \odot : G \times G \to G
> \end{align*}
> $$
> i.e., a function which takes two elements of $ G $ as arguments, and returns another element of $ G $.
>

> **Definition**: 
> A **group** is a set $ G $, together with a binary operation, $ \odot: G \times G \to G $ satisfactory of the following,
> * Associativity: $ ( a \odot b ) \odot c = a \odot ( b \odot c ) $ for all $ a,b,c \in G $.
> * Identity: $ \exists e \in G $ s.t. $ a \odot e = e \odot a = a $ for all $ a \in G $.
> * Inverses: for each $ a \in G, \exists a ^{-1} \in G $ s.t. $ a \odot a ^{-1} = a ^{-1} \odot a = e $.
> 
> Further to this, we say that a group is **abelian**, if $ a \odot b = b \odot a $ for all $ a,b \in G $.
>

> **Remark**: 
> We make note of the fact that the existence of an identity element in the group is also assurance of the non-emptiness of the set $ G $.
>

> **Example**: 
> An important example of a group is that of the integers modulo $ n \in \mathbb{Z}^{+} $, denoted by $ \mathbb{Z}/n \mathbb{Z} $ together with addition. It is fairly simple to check each of the definitive axioms hold.
> 
> We may also consider the set $ ( \mathbb{Z}/n \mathbb{Z} )^{\times} $ for $ n \in \mathbb{Z}^{+} $, i.e., the set of equivalence classes $ \overline{a} $ which have multiplicative inverses modulo $ n $. This set is a group with respect to the usual operation of multiplication.
>

> **Proposition**: 
> For a group $ G $ under the operation $ \odot $,
> * The identity element is unique.
> * For each $ a \in G $, $ a ^{-1}\in G $ is unique.
> * $ ( a ^{-1} )^{-1} = a $ for all $ a \in G $.
> 
> **Proof**:
> The proofs of the above statements are straightforward, and it is fairly clear that the proof of the final statement is reliant only on the proof of the penultimate, together with consideration of the definition of an inverse. Proof of the penultimate statement is obtained as follows.
> 
> Suppose that for an element $ a \in G $, there exist two inverses $ b,c \in G $. Then
> $$
> \begin{align*}
> c & = c \odot e             \\
> & = c \odot ( a \odot b ) \\
> & = ( c \odot a ) \odot b \\
> & = e \odot b = b
> \end{align*}
> $$
> <div class="text-right" >Q.E.D.</div>
>
