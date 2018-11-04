function [solution,optimalobjfunct] = ericMinimizef(a,b,h)
%UNTITLED Summary of this function goes here
%   Detailed explanation goes here
optimalobjfunct = b*b;
currentX = a;
solution = a;

while currentX < b
    
    currentY = (currentX - 1)^2 * sin(currentX);
    
    if currentY < optimalobjfunct
        solution = currentX;
        optimalobjfunct = currentY;
    end
    currentX = currentX + h;
    
end
end

