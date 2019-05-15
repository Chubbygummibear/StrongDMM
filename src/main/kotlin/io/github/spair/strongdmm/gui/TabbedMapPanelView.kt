package io.github.spair.strongdmm.gui

import io.github.spair.strongdmm.gui.mapcanvas.MapCanvasView
import java.awt.BorderLayout
import javax.swing.JComponent
import javax.swing.JPanel

object TabbedMapPanelView : View {
    override fun initComponent(): JComponent {
        return JPanel(BorderLayout()).apply {
            add(MapCanvasView.initComponent())
        }
    }
}
