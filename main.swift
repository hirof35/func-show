import SwiftUI

struct TrigonometricSimulatorView: View {
    @State private var degrees: Double = 0.0
    
    // 定数定義
    private let center = CGPoint(x: 200, y: 200)
    private let radius: CGFloat = 150
    
    var cosValue: Double { cos(degrees * .pi / 180.0) }
    var sinValue: Double { sin(degrees * .pi / 180.0) }
    
    var targetPoint: CGPoint {
        CGPoint(
            x: center.x + CGFloat(cosValue) * radius,
            y: center.y - CGFloat(sinValue) * radius // iOSの座標系（上がマイナス）に対応
        )
    }
    
    var body: some View {
        VStack(spacing: 0) {
            // 描画エリア
            ZStack {
                Color.white
                
                // グリッド線
                Path { path in
                    for i in stride(from: 0, to: 400, by: 10) {
                        path.move(to: CGPoint(x: i, y: 0))
                        path.addLine(to: CGPoint(x: i, y: 400))
                        path.move(to: CGPoint(x: 0, y: i))
                        path.addLine(to: CGPoint(x: 400, y: i))
                    }
                }
                .stroke(Color(red: 0.88, green: 0.88, blue: 0.88), lineWidth: 0.5)
                
                // 中心軸
                Path { path in
                    path.move(to: CGPoint(x: 0, y: 200))
                    path.addLine(to: CGPoint(x: 400, y: 200))
                    path.move(to: CGPoint(x: 200, y: 0))
                    path.addLine(to: CGPoint(x: 200, y: 400))
                }
                .stroke(Color.black, lineWidth: 1)
                
                // 単位円
                Circle()
                    .stroke(Color.red, lineWidth: 2)
                    .frame(width: radius * 2, height: radius * 2)
                    .position(center)
                
                // 各種インジケータ線の描画
                Group {
                    // cos線 (X軸への射影)
                    Path { p in
                        p.move(to: targetPoint)
                        p.addLine(to: CGPoint(x: targetPoint.x, y: center.y))
                    }.stroke(Color(red: 0, green: 0, blue: 0.75), lineWidth: 2)
                    
                    // sin線 (Y軸への射影)
                    Path { p in
                        p.move(to: targetPoint)
                        p.addLine(to: CGPoint(x: center.x, y: targetPoint.y))
                    }.stroke(Color(red: 0, green: 0.75, blue: 0), lineWidth: 2)
                    
                    // 半径線
                    Path { p in
                        p.move(to: center)
                        p.addLine(to: targetPoint)
                    }.stroke(Color(red: 0.75, green: 0, blue: 0), lineWidth: 2)
                }
                
                // 数値テキスト表示
                Text(String(format: "cos:%.3f", cosValue))
                    .font(.system(size: 14, design: .monospaced))
                    .foregroundColor(Color(red: 0, green: 0, blue: 0.75))
                    .position(x: targetPoint.x + 35, y: center.y + 10)
                
                Text(String(format: "sin:%.3f", sinValue))
                    .font(.system(size: 14, design: .monospaced))
                    .foregroundColor(Color(red: 0, green: 0.75, blue: 0))
                    .position(x: center.x + 35, y: targetPoint.y - 10)
            }
            .frame(width: 400, height: 400)
            
            // コントロールエリア（スライダ本体）
            VStack {
                Slider(value: $degrees, in: 0...360)
                    .padding(.horizontal, 20)
                Text(String(format: "Angle: %.1f°", degrees))
                    .font(.caption)
                    .foregroundColor(.gray)
            }
            .frame(width: 400, height: 50)
            .background(Color(red: 0.9, green: 0.9, blue: 0.9))
        }
        .frame(width: 400, height: 450)
    }
}
