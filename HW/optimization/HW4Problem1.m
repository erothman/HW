C = [23,1,-16,-1,52,-6,-12]
b = [62;62;3]
A = [-6,-5,25,3,-85,4,30;24,-2,28,6,-55,1,-9;9,-5,11,2,-55,-1,19]
rightSide = [0;b]
leftSide = [1;0;0;0]
middle = [-C; A]
Tableau = [leftSide,middle,rightSide]

fprintf("This is the pretableau\n")
fprintf("The starting Basis we will work with is columns 1,3,6\n")
format short
Tableau(2,:) = Tableau(2,:)/(Tableau(2,2))
Tableau(1,:) = Tableau(1,:)-(Tableau(2,:)*Tableau(1,2))
Tableau(3,:) = Tableau(3,:)-(Tableau(2,:)*Tableau(3,2))
Tableau(4,:) = Tableau(4,:)-(Tableau(2,:)*Tableau(4,2))
fprintf("Reduce the first column to rref\n")
Tableau(3,:) = Tableau(3,:)/(Tableau(3,4))
Tableau(1,:) = Tableau(1,:)-(Tableau(3,:)*Tableau(1,4))
Tableau(2,:) = Tableau(2,:)-(Tableau(3,:)*Tableau(2,4))
Tableau(4,:) = Tableau(4,:)-(Tableau(3,:)*Tableau(4,4))
fprintf("Reduce the 3rd column to rref\n")
Tableau(4,:) = Tableau(4,:)/(Tableau(4,7))
Tableau(1,:) = Tableau(1,:)-(Tableau(4,:)*Tableau(1,7))
Tableau(3,:) = Tableau(3,:)-(Tableau(4,:)*Tableau(3,7))
Tableau(2,:) = Tableau(2,:)-(Tableau(4,:)*Tableau(2,7))
fprintf("Reduce the 6th column to rref\n")
fprintf("The Column that would cause the largest change is column 2:")
fprintf(" it is the one with the largest number in the top row./n")
fprintf("So we will pivot based on that column.\n")
fprintf("We can ignore pivoting on the 3rd row since that is negative.")
fprintf("The two ratios to decide the pivot are:")
Tableau(2,9)/Tableau(2,3)
Tableau(4,9)/Tableau(4,3)
fprintf("Since the ratio when pivoted on the 2nd row is smaller we shall")
fprintf("pivot there.\n")
Tableau(2,:) = Tableau(2,:)/(Tableau(2,3))
Tableau(1,:) = Tableau(1,:)-(Tableau(2,:)*Tableau(1,3))
Tableau(3,:) = Tableau(3,:)-(Tableau(2,:)*Tableau(3,3))
Tableau(4,:) = Tableau(4,:)-(Tableau(2,:)*Tableau(4,3))
fprintf("Now the only column with a positive number at the top is column")
fprintf("4.\nSo we will pivot on that column.\nSince the 4th row has a")
fprintf("negative value in that column, we will ignore that number for")
fprintf("being a pivot number.\nSo the 2 potential pivot ratios are:")
Tableau(2,9)/Tableau(2,5)
Tableau(3,9)/Tableau(3,5)
fprintf("The smaller of the ratios is the one of the 3rd row so that is")
fprintf("the one that will reach 0 first.\nSo we will pivot on the 3rd")
fprintf("row and the 4th column.\n")
Tableau(3,:) = Tableau(3,:)/(Tableau(3,5))
Tableau(1,:) = Tableau(1,:)-(Tableau(3,:)*Tableau(1,5))
Tableau(2,:) = Tableau(2,:)-(Tableau(3,:)*Tableau(2,5))
Tableau(4,:) = Tableau(4,:)-(Tableau(3,:)*Tableau(4,5))
fprintf("Since there is no positive number on the top row that represents ")
fprintf("a variable, we are done.\nIf we pivot on any row the objective ")
fprintf("function value will just go up.\n")
fprintf("So the value for the objective function is -68 which occurs at ")
fprintf("the values of:\n")
fprintf("[x1    [0\n x2     1\n x3     0\n x4  =  9\n x5     0\n x6     10\n x7]    0]\n")