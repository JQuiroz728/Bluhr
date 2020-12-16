# Bluhr
Transform your images into geometric primitives using this web application

Using the [primitive](https://github.com/fogleman/primitive) library, for each unique mode, the algorithm finds the most optimal shape that can be drawn to minimize the error between the original image and the resulting image. Results shown utilize 100 distinct shapes. More shapes will produce results resembling a closer appearance to the original photo, at the cost of a slower rendering speed.

Original Photo: My dog Beckham 
![](chavy.jpg =100x20)

Results (100 distinct shapes using various modes): 
1. Triangles           
2. Rectangles         
3. Ellipses               
4. Circles          
5. Rotated Rectangles
![](results1.png)

6. Bézier curves
7. Rotated Ellipses
8. Polygons
9. Combination of all
![](results2.png)
