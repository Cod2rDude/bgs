package vector2

import "math"

// Public Types
type Vector2 struct {
    X float64
    Y float64
}

// Public Constructors
func New(x, y float64) Vector2 {
    return Vector2{
        X: x,
        Y: y,
    }
}

func FromAngle(radians float64) Vector2 {
    return Vector2{
        X: math.Cos(radians),
        Y: math.Sin(radians),
    }
}

func Zero() Vector2 {
    return Vector2{0, 0}
}

func One() Vector2 {
    return Vector2{1, 1}
}

// Public Methods
func (v Vector2) Magnitude() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector2) Normalize() Vector2 {
    mag := v.Magnitude()
    if mag == 0 {
        return Vector2{0, 0}
    }
    return Vector2{
        X: v.X / mag,
        Y: v.Y / mag,
    }
}

func (v Vector2) Dot(other Vector2) float64 {
    return v.X*other.X + v.Y*other.Y
}

func (v Vector2) Cross(other Vector2) float64 {
    return v.X*other.Y - v.Y*other.X
}

func (v Vector2) Distance(other Vector2) float64 {
    dx := v.X - other.X
    dy := v.Y - other.Y
    return math.Sqrt(dx*dx + dy*dy)
}

func (v Vector2) Lerp(other Vector2, t float64) Vector2 {
    if t < 0.0 {
        t = 0.0
    } else if t > 1.0 {
        t = 1.0
    }

    return Vector2{
        X: v.X + (other.X-v.X)*t,
        Y: v.Y + (other.Y-v.Y)*t,
    }
}

func (v Vector2) AngleBetween(other Vector2) float64 {
    dot := v.Dot(other)
    magV := v.Magnitude()
    magOther := other.Magnitude()

    if magV == 0 || magOther == 0 {
        return 0
    }

    cosTheta := dot / (magV * magOther)

    if cosTheta < -1.0 {
        cosTheta = -1.0
    } else if cosTheta > 1.0 {
        cosTheta = 1.0
    }

    return math.Acos(cosTheta)
}

func (v Vector2) Rotate(radians float64) Vector2 {
    sin := math.Sin(radians)
    cos := math.Cos(radians)

    return Vector2{
        X: v.X*cos - v.Y*sin,
        Y: v.X*sin + v.Y*cos,
    }
}

func (v Vector2) RotateAround(pivot Vector2, radians float64) Vector2 {
    shifted := v.Subtract(pivot)
    rotated := shifted.Rotate(radians)
    
    return rotated.Add(pivot)
}

func (v Vector2) ToAngle() float64 {
    return math.Atan2(v.Y, v.X)
}

func (v Vector2) LookAt(target Vector2) float64 {
    diff := target.Subtract(v)
    return math.Atan2(diff.Y, diff.X)
}

func (v Vector2) MoveTowards(target Vector2, maxDistance float64) Vector2 {
    diff := target.Subtract(v)
    dist := diff.Magnitude()

    if dist <= maxDistance || dist == 0 {
        return target
    }

    return v.Add(diff.Scale(maxDistance / dist))
}

func (v Vector2) Reflect(normal Vector2) Vector2 {
    dot := v.Dot(normal)
    return v.Subtract(normal.Scale(2 * dot))
}

func (v Vector2) Project(onto Vector2) Vector2 {
    magSq := onto.X*onto.X + onto.Y*onto.Y
    if magSq == 0 {
        return Vector2{0, 0}
    }
    dot := v.Dot(onto)
    return onto.Scale(dot / magSq)
}

func (v Vector2) Add(other Vector2) Vector2 {
    return Vector2{
        X: v.X + other.X,
        Y: v.Y + other.Y,
    }
}

func (v Vector2) Subtract(other Vector2) Vector2 {
    return Vector2{
        X: v.X - other.X,
        Y: v.Y - other.Y,
    }
}

func (v Vector2) Multiply(other Vector2) Vector2 {
    return Vector2{
        X: v.X * other.X,
        Y: v.Y * other.Y,
    }
}

func (v Vector2) Divide(other Vector2) Vector2 {
    return Vector2{
        X: v.X / other.X,
        Y: v.Y / other.Y,
    }
}

func (v Vector2) Scale(scalar float64) Vector2 {
    return Vector2{
        X: v.X * scalar,
        Y: v.Y * scalar,
    }
}