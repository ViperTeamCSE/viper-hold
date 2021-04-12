using System;

public class BadMath : IDisposable
{

    public int factorial(int x) {
        return x * factorial(x - 1);
    }

    public void Dispose()
    {
        
    }
}