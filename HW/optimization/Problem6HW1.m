x1Max = 0;
x2Max = 0;
maxValue = 0;
x1 = 0;
h = 0.001;
while x1 <= 6
    x2 = 0;
    while x2 <= 6 - x1
        if (x2 >= log(x1))
            funcValue = 1 + (x1)^2 * (x2 - 1)^3 * exp(-(x1 + x2));
            if funcValue > maxValue
                x1Max = x1;
                x2Max = x2;
                maxValue = funcValue;
            end
        end
        x2 = x2 + h;
    end
    x1 = x1 + h;
end
fprintf("maximized solution is [%f %f] with a value of %f\n", x1Max, x2Max, maxValue);