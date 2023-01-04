(ns a-01.sequences)

(defn Seq
  []
  (def colors (seq ["red" "green" "blue"]))
  (println colors)
  (println (cons "yellow" colors))
  (println (cons colors "black"))
  (println (conj colors "yellow"))
  (println (conj ["red" "green" "blue"] "yellow"))
  (println (concat colors (seq ["black" "white"])))
  (println (distinct (seq [1 2 1 3 3 4 4 4 5])))
  (println (reverse colors))
  (println (first colors))
  (println (rest colors))
  (println (last colors))
  (println (sort (seq [1 3 2 5 4])))
  )

(Seq)
